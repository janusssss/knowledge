### 入门案例

```bash
# 版本Scala 2.13-3.7.1
# 启动zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
# 启动kafka
bin/kafka-server-start.sh config/server.properties
# 创建topic
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic test

old /data1/hxxt_yx/kafka/bin/kafka-topics.sh --create --zookeeper 10.142.149.168:2186 --replication-factor 1 --partitions 3 --topic janus

# 查看topic
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
old /data1/hxxt_yx/kafka/bin/kafka-topics.sh --list --zookeeper  10.142.149.168:2186
# 创建一个consumer
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
# 创建一个producer
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test

# 查看所有消费者组
kafka-consumer-groups.sh --bootstrap-server <broker_address> --list
```

### 配置

```bash
# 保留时间
log.retention.hours=168
# 保留大小
log.retention.bytes=1073741824
```

### 概述

​	Apache Kafka 是一个分布式流处理平台，最初被设计用来统一处理大型企业中的所有实时数据流。它兼具**高吞吐量**、**低延迟**、**可扩展性**、**持久化**和**容错性**，使其适用于从实时日志聚合、网站活动追踪到大规模事件流处理等多种场景。 

**核心设计思想与特点：** 

1. **日志即核心 (Log-Centric Design)**：Kafka 将数据流存储在**分区**（Partition）中，每个分区是一个**持久化、有序**的日志文件。这种设计不同于传统的消息队列，更像数据库日志，为高效存储和处理提供了基础。
2. **拥抱文件系统 (Embrace the Filesystem)**：Kafka 重度依赖文件系统进行存储和缓存，利用现代操作系统的**页缓存**（Page Cache）和优化的磁盘 I/O（如 `sendfile` 系统调用）来实现接近网络速度的性能，而非完全依赖内存。这使得它可以处理**近乎无限**的数据量而不牺牲性能。
3. **常量时间操作 (Constant Time Operations)**：基于追加写入和读取日志的结构，使得读写操作的时间复杂度为 O(1)，性能与数据量大小解耦。
4. 

 **高效性 (Efficiency)**：    *   **批处理 (Batching)**：通过将消息分组（消息集）进行网络传输和磁盘 I/O，显著提高吞吐量。    *   **零拷贝 (Zero-Copy)**：利用 `sendfile` 等技术减少数据在内核态和用户态之间的拷贝，提高效率。    *   **端到端压缩 (End-to-End Compression)**：支持对消息批次进行压缩，优化网络带宽使用。 5.  **分区与水平扩展 (Partitioning & Horizontal Scaling)**：主题（Topic）被划分为多个分区，可以分布在整个集群中。这不仅实现了**水平扩展**，还支持**并行处理**。每个分区可以有多个副本。 6.  **复制与容错 (Replication & Fault Tolerance)**：每个分区可以被复制到多个 broker 上，形成一个**领导者**（Leader）和多个**追随者**（Follower）。通过维护**同步副本集**（ISR），Kafka 能够在 broker 故障时自动进行故障转移，保证数据的可用性和持久性（在至少一个 ISR 副本存活的前提下）。 7.  **消费者模型 (Consumer Model)**：消费者采用**拉取**（Pull）模式从 broker 获取数据，可以自由控制消费的位置（偏移量 Offset），支持重复消费或跳过数据。消费者按**消费者组**（Consumer Group）进行管理，组内每个分区同一时间只被一个消费者消费。 8.  **多种消息传递语义 (Delivery Semantics)**：支持**最多一次**、**至少一次**，并通过**幂等性**（Idempotence）和**事务**（Transactions）支持**精确一次**（Exactly-Once）语义，满足不同场景对数据一致性的要求。 9.  **日志压缩 (Log Compaction)**：可选功能，确保每个消息键（Key）的最新值在日志中至少保留一份，适用于需要维护最新状态的场景（如数据库变更日志）。 10. **配额管理 (Quotas)**：可以对客户端的网络带宽和请求速率进行限制，防止个别客户端耗尽集群资源，保障多租户环境下的稳定性和公平性。 

总而言之，Kafka 通过其独特的日志中心化设计、对文件系统和操作系统特性的深度利用、高效的批处理与复制机制，构建了一个能够处理海量实时数据流、具备高吞吐、低延迟、强持久性和高可用性的平台。它不仅是一个消息队列，更是一个强大的**流存储**和**流处理**基础设施工具。 

 

 
 

 
    Replication（副本机制）：这是 Kafka 高可用性的核心。

    Rebalance（重平衡）：这是 Kafka 消费者组的核心行为。

我会为您详细解释这两个机制，因为它们都是 Kafka 架构中至关重要的部分。
1. Kafka Replication（副本机制）

这是保证 Kafka 在部分服务器故障时仍能正常工作的 “备份”机制。

核心思想： 将每个分区的数据复制到多个 Broker（Kafka 服务器）上，从而提供数据冗余和故障转移。

关键概念：

    Leader副本：每个分区都有一个 Leader。所有的生产者写入和消费者读取请求都只与 Leader 副本交互。

    Follower副本：Follower 副本的任务就是从 Leader 副本异步地拉取数据，保持与 Leader 的数据同步。

    ISR（In-Sync Replica，同步副本集合）：所有与 Leader 保持同步（不仅仅是存活，而且数据差距在可接受范围内）的 Follower 副本和 Leader 副本自己组成的集合。

工作机制：

    生产者发送消息到某个 Topic 的 Partition Leader。

    Leader 将消息写入其本地日志。

    所有 ISR 列表中的 Follower 副本会从 Leader 拉取消息，并回复一个 ACK。

    当 Leader 收到所有 ISR 的 ACK 后，才认为这条消息已提交。

    生产者可以配置等待的 ACK 级别（acks配置）：

        acks=0：不等待，性能最高，数据可能丢失。

        acks=1：只等待 Leader 确认（默认）。

        acks=all：等待所有 ISR 副本都确认，最安全，性能最低。

故障转移：
如果 Leader 副本宕机了，Kafka 会从 ISR 中自动选举一个新的 Leader。由于新 Leader 拥有最新的数据，这个过程对生产者和消费者是透明的，确保了高可用性。

所以，您说的“备份机制”，其核心就是 Replication（副本机制）。
2. Kafka Rebalance（重平衡机制）

这是 Kafka 消费者组（Consumer Group） 的核心行为，目的是在消费者成员发生变化时，重新分配分区（Partition）的所有权。

触发 Rebalance 的条件：

    消费者加入群组：例如，你启动了一个新的消费者实例。

    消费者离开群组：例如，消费者实例崩溃、被关闭、或者长时间心跳超时。

    订阅的 Topic 发生变化：例如，管理员在运行时创建了一个新的分区。

    消费者取消订阅某个 Topic。

Rebalance 的目标：
保证一个分区在同一个时刻只被同一个消费者组内的一个消费者消费，同时尽可能均衡地将所有分区分配给组内的所有消费者。

举个例子：
假设一个 Topic 有 3 个分区（P0, P1, P2），一个消费者组（Group A）最初有 2 个消费者（C1, C2）。

    初始分配：C1 消费 P0 和 P1，C2 消费 P2。

    触发 Rebalance：现在组内加入了一个新的消费者 C3。

    重平衡后：Kafka 会重新分配，可能的结果是 C1 消费 P0，C2 消费 P1，C3 消费 P2。这样每个消费者的负载更加均衡。

Rebalance 的代价：
在 Rebalance 期间，整个消费者组会停止工作（Stop-the-World），直到分配完成。这段时间内没有消费者能够处理消息，可能会造成消费延迟。

改进：Kafka 的 Cooperative Sticky 策略
早期版本的 Rebalance 策略（如 Eager）会放弃所有当前的分区分配，然后重新分配。而新的 Cooperative Sticky 策略是一种“合作式”的重平衡，它只重新分配必要的那部分分区，最大限度地减少“停止世界”的影响。



同一消费者组内：一个分区只能被一个消费者消费
不能并发：同一个分区的消息必须按顺序被单个消费者处理
