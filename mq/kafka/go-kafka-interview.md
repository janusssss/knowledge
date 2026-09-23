# Go + Kafka 面试常问问题

>

## 一、基础与消息模型

1. **Kafka 的架构里，Topic / Partition / Replica / ISR / Leader 分别是什么关系？**
   一个 Topic 分成多个 Partition，每个 Partition 是一个有序、不可变、只追加的日志；每个 Partition 有 1 个 Leader 和 N-1 个 Follower；ISR（In-Sync Replica）是当前与 Leader 保持同步的副本集合，只有 ISR 里的副本才有资格被选为新 Leader。

2. **为什么 Kafka 能把吞吐做到这么高？**
   - 顺序写磁盘 + Page Cache，几乎不做随机 IO
   - 零拷贝（`sendfile`）把数据直接从 Page Cache 送到网卡，绕过用户态
   - 批量（batch）+ 压缩（压缩以 batch 为单位，压缩比更高）
   - Partition 级并行，生产消费两端都能水平扩展
   - 稀疏索引（`.index` 是 offset 索引，`.timeindex` 是时间戳索引），查 offset 是二分查找

3. **Partition 数量怎么定？多了少了各有什么问题？**
   - 少：并行度不够，消费端上不去
   - 多：Controller 元数据变大、故障恢复（Leader 选举）变慢、端到端延迟升高、每个分区都占文件句柄和内存
   - 经验值：以吞吐目标反推，`目标吞吐 / 单分区吞吐`；一般单分区 10 MB/s 级别，同时别超过消费者线程数的合理倍数
   - **Kafka 不支持减少分区**，只能增加，增加会打乱 key 的哈希分布，所以一开始宁可多别多太多

4. **同一个 key 的消息为什么一定在同一个 Partition？**
   因为 Go 客户端默认用 `hash(key) % numPartitions` 选分区（`kafka.Hash`），同 key 必然落同一分区，从而保证分区内有序。注意：**只在单分区内有序**，跨分区不保证全局有序。

5. **什么是幂等 Producer 和事务？分别解决什么问题？**
   - 幂等（`enable.idempotence=true`）：给每条消息带 ProducerID + SequenceNumber，Broker 端去重，解决重试导致的重复投递。**只保证单分区、单会话内幂等**
   - 事务（`transactional.id`）：跨分区、跨会话的原子写，解决「读-处理-写」的 exactly-once（配合 `isolation.level=read_committed` 的消费者）

6. **acks 的三个取值意味着什么？**
   | 值 | 含义 | 风险 |
   |---|---|---|
   | 0 | 发出去就不管 | 丢消息无法感知 |
   | 1 | Leader 写完就返回 | Leader 挂且未同步则丢 |
   | -1/all | ISR 全部写完才返回 | 最安全，延迟最高 |
   配合 `min.insync.replicas=2` 才能真正防丢。

7. **什么是 LEO、HW、Committed Offset？**
   - LEO（Log End Offset）：副本日志末尾下一条消息的 offset
   - HW（High Watermark）：ISR 中所有副本都已同步到的最小 LEO，消费者只能读到 HW 之前的数据
   - Committed Offset：消费者组已确认消费的位置，存在 `__consumer_offsets` 里

## 二、Producer 调优

8. **`batch.size`、`linger.ms`、`compression.type` 三者怎么联动？**
   消息先在 RecordAccumulator 里按分区攒批，攒满 `batch.size` 或等够 `linger.ms` 就发出去。加大 `batch.size` + 适度 `linger.ms`（5~100ms）+ 用 `lz4`/`zstd` 压缩，是提吞吐最有效的一招，代价是延迟上升。

9. **`max.in.flight.requests.per.connection` 和顺序性有什么关系？**
   关掉幂等时，该值 > 1 会因为重试导致乱序。开了幂等（并且 `max.in.flight <= 5`）Broker 会按序号重排，既保顺序又保吞吐。

10. **Producer 发消息的完整流程？**
    `send()` → 序列化 → 分区器选分区 → 累加器按分区攒批 → Sender 线程拉到批次 → 按 Broker 分组请求 → 网络发送 → 收到响应后回调（`kafka.Producer` 在 Sarama 里对应 `AsyncProducer`/`SyncProducer` 的这个链条）。

11. **`retries`、`delivery.timeout.ms`、`request.timeout.ms` 的关系？**
    `delivery.timeout.ms >= linger.ms + request.timeout.ms`，它约束「一条消息从进累加器到最终失败」的总时长，超时后回调报错。设太短反而更容易丢。

12. **消息发不进去怎么排查？**
    - `NotLeaderForPartition`：元数据过期，客户端会自动刷新并重试
    - `RecordTooLarge`：撞 `message.max.bytes` / `max.request.size`
    - `TimeoutException`：网络或 Broker 压力，看 `linger.ms` 与重试配置
    - `QueueFull`：`buffer.memory` 满了，说明生产速度远大于发送速度

## 三、Consumer 与消费组

13. **Rebalance 是怎么触发的？有哪些协议？**
    触发：成员加入/离开/心跳超时/订阅关系变化/分区数变化。
    协议演进：
    - **Eager（Range/RoundRobin）**：全部撤销再重新分配，STW
    - **Cooperative Sticky**：增量协作式，只撤销真正需要移动的分区，避免全组重平衡
    - **KIP-848（Broker 端 rebalance，新协议）**：由 Group Coordinator 计算，客户端不再主导

14. **`session.timeout.ms`、`heartbeat.interval.ms`、`max.poll.interval.ms` 分别是什么？**
    - `heartbeat.interval.ms`：心跳间隔，独立线程发送
    - `session.timeout.ms`：多久收不到心跳就认为成员死了
    - `max.poll.interval.ms`：两次 `Poll()` 之间最长间隔，**超过就被踢出组**，这是「消费逻辑太慢导致反复 rebalance」的元凶
    关系：`max.poll.interval.ms > 单批处理耗时`，`session.timeout.ms` 一般取 `3 * heartbeat.interval.ms`。

15. **怎么实现 at-least-once / at-most-once / exactly-once？**
    - at-most-once：先提交 offset 再处理，挂了就丢
    - at-least-once：先处理再提交，挂了会重复（**最常见的生产选择**，靠业务幂等兜底）
    - exactly-once：用事务把「处理结果 + offset 提交」绑在一起（`SendOffsetsToTransaction`）

16. **手动提交 offset 有几种方式？各有什么坑？**
    - `CommitSync`：阻塞、可靠、慢
    - `CommitAsync`：不阻塞、可能失败、**回调里的重试会覆盖更新的 offset**，一般配合「关闭前最后同步提交一次」
    - 用 `MarkMessage` / `CommitOffsets` 时注意提交的是 `offset + 1`
    - **先处理再提交**，且提交失败要能重试

17. **一个消费者组里消费者数量大于分区数会怎样？**
    多出来的消费者空转（拿不到分区）。所以 **并行度上限 = 分区数**。这是设计时必须同步考虑的点。

18. **消费者 lag 怎么看、怎么处理？**
    - 指标：`records-lag-max`、`kafka-consumer-groups.sh --describe` 里的 LAG 列
    - 短期突增（流量峰值）+ 处理变慢（下游 DB 慢）是两大来源
    - 手段：加分区 + 加消费者、优化处理逻辑（批量写库、去掉同步 RPC）、加大 `fetch.min.bytes` / `max.partition.fetch.bytes`、把慢逻辑异步化

19. **`auto.offset.reset` 的 earliest / latest / none 区别？**
    只在「没有已提交 offset」时生效（新组或 offset 过期）。`none` 会直接抛异常，适合要求显式初始化的场景。

20. **消费顺序性怎么保证？**
    生产端用 key 把相关消息路由到同一分区；消费端**单分区单 Goroutine 串行处理**。Go 里常见做法是每个分区一个 worker + 分区级队列，而不是全局并发处理后再乱序写。

## 四、Go 客户端（sarama / kafka-go / confluent-kafka-go）

21. **主流 Go Kafka 客户端怎么选？**
    | 库 | 特点 | 适合 |
    |---|---|---|
    | `IBM/sarama`（原 Shopify/sarama） | 纯 Go、生态最广、API 啰嗦、维护一度停滞 | 通用场景 |
    | `segmentio/kafka-go` | API 友好、`Reader`/`Writer` 抽象好、内置消费组 | 中小项目、快速开发 |
    | `confluent-kafka-go` | 包 librdkafka（CGO），功能最全、性能最好、事务支持最完整 | 对性能和 exactly-once 有要求 |
    | `twmb/franz-go` | 纯 Go、现代、性能强、功能覆盖全（含事务、KIP-848） | 新项目首选 |

22. **Sarama 的 `AsyncProducer` / `SyncProducer` / `ConsumerGroup` 怎么用？**
    - `SyncProducer`：每条 `SendMessage` 阻塞等 ack，简单但吞吐低
    - `AsyncProducer`：走 `Input()` channel + `Successes()`/`Errors()` 两个输出 channel，**必须同时消费这两个 channel，否则会阻塞整个 producer**
    - `ConsumerGroup`：实现 `ConsumeClaim(session, claim)`，**`claim.Messages()` 的 channel 关掉前不要 return**，否则会触发 rebalance

23. **消费组的 `ConsumerGroupHandler` 三个方法分别在什么时候调用？**
    ```
    Setup(session)     // rebalance 后、消费开始前，做初始化
    ConsumeClaim(...)  // 真正的消费循环，每个 claim 一个 Goroutine
    Cleanup(session)   // 所有 ConsumeClaim 返回后，做收尾/最后提交
    ```
    典型 bug：在 `Setup` 里起 Goroutine 然后 `Cleanup` 里没停；或者 `ConsumeClaim` 里 `return nil` 太早导致分区被回收。

24. **Sarama 里怎么优雅关闭？**
    ```go
    // 先 CloseAsyncProducer，把 input channel 关掉
    producer.Close()      // 会 flush 未发完的批次
    consumer.Close()      // 会触发一次 rebalance 和最后一次 offset 提交
    ```
    直接用 `CloseAsyncProducer` 后再 `Close`；别忘了 `sarama.NewConfig()` 里 `config.Producer.Return.Successes = true`，否则 `Successes()` channel 永远没数据。

25. **消费时怎么防止 Goroutine 泄漏？**
    - 用 `context.Context` 贯通，`select` 里同时监听 `ctx.Done()`
    - 每个分区的处理 Goroutine 用 `sync.WaitGroup` 等待
    - 别在 handler 里无脑 `go func(){}` 处理消息，会打乱 offset 提交语义

26. **消息体怎么序列化？**
    常用 `JSON`、`Protobuf`、`Avro`。生产建议 Schema Registry + Avro/Protobuf，好处是兼容性校验（向前/向后兼容）。**JSON 没有 schema 演进约束，字段改名就炸**。

## 五、可靠性、监控与运维

27. **怎么保证消息不丢？**
    三个环节都要看：
    - 生产端：`acks=all` + `retries>0` + 处理 Error channel + 幂等
    - Broker 端：`replication.factor>=3` + `min.insync.replicas>=2` + `unclean.leader.election.enable=false`
    - 消费端：关闭自动提交，处理成功后再提交

28. **怎么保证不重复？**
    Kafka 本身在 at-least-once 下无法避免重复，只能业务侧幂等：唯一键去重表、Redis SETNX、状态机（只允许特定状态迁移）。或者上事务做 exactly-once。

29. **`unclean.leader.election.enable` 开着有什么后果？**
    允许非 ISR 副本当 Leader，可用性提升但**会丢数据**。金融类场景必须 false。

30. **日志保留策略有哪些？**
    `retention.ms` / `retention.bytes` / `log.segment.bytes` / `log.retention.check.interval.ms`，还有 `cleanup.policy=compact`（压缩，只保留每个 key 的最新值，适合 CDC、状态快照）。

31. **关键监控指标有哪些？**
    - Broker：`UnderReplicatedPartitions`、`OfflinePartitionsCount`、`ActiveControllerCount`（**必须为 1**）、`RequestHandlerAvgIdlePercent`、磁盘使用率
    - Producer：`record-error-rate`、`request-latency-avg`、`batch-size-avg`、`buffer-available-bytes`
    - Consumer：`records-lag-max`、`fetch-rate`、`commit-latency-avg`、`rebalance-rate-per-hour`
    - Go 端：把 sarama 的 `Metrics` 接到 Prometheus，或直接采 JMX exporter

32. **消息积压了几百万条，怎么紧急处理？**
    1. 先判断是「消费端变慢」还是「生产端暴涨」
    2. 临时扩容消费者（**不超过分区数**），必要时先把消息**转存到新 Topic**（新建更多分区的 Topic 做接力）
    3. 紧急情况下先批量落库/落盘，恢复 offset，再离线补处理
    4. 事后补告警阈值

33. **Controller 是干什么的？脑裂怎么办？**
    Controller 负责分区 Leader 选举、ISR 变更、Topic 增删。通过 ZK（老版本）或 KRaft（新版本）选主。`ActiveControllerCount != 1` 就说明集群脑裂或元数据异常。

34. **KRaft 和 ZooKeeper 的区别？**
    KRaft 用 Raft 协议自己管元数据，去掉了 ZK 依赖，元数据变更更快、支持更多分区（百万级）。Kafka 3.3+ 生产可用，4.0 开始已完全移除 ZK。

## 六、场景设计题（面试最爱）

35. **设计一个订单状态变更的实时同步系统，要求顺序、不丢、可回放。**
    - 用 `orderId` 作 key，保证同订单落同一分区
    - 分区数按 QPS 和峰值算，预留 2x
    - 消费端单分区串行 + 落库幂等（订单号 + 状态版本号）
    - `retention` 设 7 天以上，支持从任意 offset 回放重建

36. **如何用 Kafka 做延迟队列？**
    - 方案 A：多级 Topic + 定时搬运（比如 1min/5min/30min 三级）
    - 方案 B：单 Topic 存到期时间戳，消费端 `Pause` 分区 + 到点再 `Resume`（注意 `max.poll.interval.ms`）
    - 方案 C：Kafka + 时间轮 / 定时任务扫描（最简单，量小首选）

37. **Kafka 和 RabbitMQ / RocketMQ / Pulsar 怎么选？**
    - Kafka：日志流、大数据管道、高吞吐、可回放
    - RabbitMQ：复杂路由、低延迟、消息级 ack，吞吐一般
    - RocketMQ：事务消息、延迟消息、顺序消息场景好，国内生态强
    - Pulsar：存储计算分离、多租户、支持队列+流，运维复杂度高

38. **Kafka 为什么不支持「消息级 TTL」而只有 Topic 级？**
    因为 Partition 是只追加的有序日志，删除中间某条消息会破坏 offset 连续性。只能用 compact（key 级别清理）或按时间/大小整体删除 segment 来近似。

## 七、Go 常见代码题

39. **写一个保证不丢的生产者骨架：**
    ```go
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true
    config.Producer.Return.Errors = true
    config.Producer.Idempotent = true
    config.Net.MaxOpenRequests = 1 // 幂等要求
    config.Version = sarama.V3_6_0_0

    producer, err := sarama.NewAsyncProducer(brokers, config)
    // 必须把两个输出 channel 消费掉，否则阻塞
    go func() {
        for range producer.Successes() {}
    }()
    go func() {
        for err := range producer.Errors() {
            log.Printf("send failed: %v", err)
        }
    }()
    ```

40. **写一个带优雅退出的消费组 handler：**
    ```go
    func (h *handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
        for {
            select {
            case msg, ok := <-claim.Messages():
                if !ok {
                    return nil // 分区被撤销，正常退出
                }
                if err := h.process(msg); err != nil {
                    return fmt.Errorf("process: %w", err)
                }
                sess.MarkMessage(msg, "")
            case <-sess.Context().Done():
                return nil
            }
        }
    }
    ```

41. **为什么 `sess.Context()` 也要监听？**
    rebalance 或 `Close()` 时会 cancel 它，只在 `claim.Messages()` 上等可能永远等不到关闭信号，导致 Goroutine 卡住。

42. **Sarama 里 `Consumer.Group.Rebalance.GroupStrategies` 怎么配？**
    ```go
    config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
        sarama.NewBalanceStrategySticky(),
    }
    config.Consumer.Group.Rebalance.Timeout = 60 * time.Second
    ```
    从 Range 换成 Sticky/Cooperative 能显著减少 rebalance 期间的分区抖动。

43. **批量处理时怎么控制提交频率？**
    攒够 N 条或 T 毫秒后统一 `sess.Commit()`，别每条都提交（会有大量 `__consumer_offsets` 写入），也别攒太久（失败重放代价大）。

44. **`kafka-go` 的 `Reader` 怎么配？**
    ```go
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:        []string{"localhost:9092"},
        GroupID:        "order-consumer",
        Topic:          "orders",
        MinBytes:       1e4,      // 10KB
        MaxBytes:       1e6,      // 1MB
        MaxWait:        time.Second,
        CommitInterval: time.Second,
        StartOffset:    kafka.FirstOffset,
    })
    ```
    注意 `CommitInterval` 是异步提交周期，要严格 at-least-once 就设 0 然后手动 `CommitMessages`。

## 八、易错点速查

- 提交 offset 时要 `msg.Offset + 1`
- 消费者数 > 分区数 → 空转
- `max.poll.interval.ms` 太小 + 处理慢 → 无限 rebalance
- `Successes()` 不消费 → 生产者静默阻塞
- 幂等 Producer 必须 `acks=all`、`retries>0`、`max.in.flight<=5`
- 加分区会破坏 key 的哈希分布，慎用
- `min.insync.replicas` 没配，`acks=all` 也可能丢
- Topic 名大小写敏感，`__consumer_offsets` 别乱删
