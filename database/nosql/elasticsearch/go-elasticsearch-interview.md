# Go + Elasticsearch 面试常问问题

>

## 一、底层原理

1. **倒排索引到底是什么结构，为什么快？**
   核心是 `term -> posting list`（文档 ID 列表）的映射。写入时经过分词（analysis）把文本拆成 term，建立倒排；查询时按 term 直接定位 posting list，做交集/并集。相比全表扫描的 O(N)，这是近似 O(1) 的定位。
   底层实现是 FST（有限状态转换器）+ 跳表（skip list）：FST 共享 term 前后缀省内存，posting list 用跳表支撑多条件合并（skip 加速求交）。

2. **ES 和 Lucene 什么关系？**
   一个 ES 索引 = 多个 shard，**每个 shard 就是一个 Lucene 索引**。ES 负责分布式（路由、复制、聚合归并、集群状态），Lucene 负责单机检索。

3. **Doc Values 和 Fielddata 的区别？**
   | | Doc Values | Fielddata |
   |---|---|---|
   | 存储 | 磁盘列存储（正排） | 堆内存 |
   | 默认开启 | 除 text 外都开 | 默认关闭 |
   | 用途 | 排序、聚合、script | 对 text 做排序聚合（不推荐） |
   `fielddata: true` 在 text 字段上打开的第二天就会 OOM，正确做法是用 `keyword` 子字段。

4. **`_source`、`store`、`_all` 分别是什么？**
   - `_source`：存原始 JSON，默认开启，用于返回文档和 reindex
   - `store`：单独把某字段存成可独立取出的形式，用于「只想取这一个字段又不想开 `_source`」的极端省空间场景
   - `_all`：7.x 起已移除，等价于 `copy_to` 到自定义字段

5. **一个索引的真实磁盘构成？**
   segment（不可变倒排）+ commit point + translog + `.del` 删除标记。写入是 append 新 segment，删除只是打标记，靠 merge 真正清理。

6. **为什么 segment 不可变？**
   不可变使得已写入的 segment 无需加锁、可被 OS Page Cache 缓存、可以被随意复制。代价是更新 = 删旧 + 写新，以及随之而来的 merge 压力。

7. **refresh / flush / merge 分别做什么？**
   - `refresh`：把内存 buffer 里的文档写成新的 segment 并打开，**使文档可被搜索**。默认 1s，这就是「ES 是准实时」的原因
   - `flush`：触发 Lucene commit（fsync segment 目录）+ 清空 translog，默认 30min 或 translog 达 512MB
   - `merge`：把多个小 segment 合并成大 segment，同时真正删除已标记删除的文档。发生在后台，是 IO 大户

8. **translog 的作用是什么？**
   文档进 buffer 时同步写 translog（默认 `request` 级别 durability 是异步，`index.translog.durability=request` 则每次请求都 fsync）。进程崩溃后可以用 translog 恢复未 commit 的数据。

9. **shard 的路由规则，为什么不能随便改分片数？**
   `shard = hash(routing) % number_of_primary_shards`。分片数变了哈希结果全变，已有的文档就找不到了。所以**分片数创建后不可改**（除 `_split`/`_shrink` 这类需重建的操作）。

10. **集群健康色代表什么？**
    - green：全部主副本都分配好了
    - yellow：主分片正常，副本未分配（单节点集群常态）
    - red：有主分片未分配，**会导致部分数据不可读**
    注意 red 不等于集群不能用，是「有分片挂了」。

11. **脑裂（split brain）怎么避免？**
    设置 `discovery.seed_hosts` + `cluster.initial_master_nodes`，并使用 `node.roles: [master]` 的专用 master 节点（至少 3 个）。7.x 起 ES 自己实现了类 Raft 的选举，`minimum_master_nodes` 已移除。

## 二、Mapping 与建模

12. **`text` 和 `keyword` 的区别，怎么选？**
    - `text`：会分词，用于全文检索（`match` 查询）。不能排序聚合
    - `keyword`：整体作为一个 term，用于精确匹配、排序、聚合（`term` 查询）
    实战写法：一个字段两个子字段，`name` 用 text 检索，`name.keyword` 用 keyword 排序聚合。

13. **动态映射（dynamic mapping）有什么坑？**
    - 数字字符串 `"123"` 会被识别成 long，导致 `"00123"` 丢失前导零
    - 长数字时间戳可能被当 date
    - 新字段自动膨胀，产生 **mapping explosion**（字段数上万后集群状态爆大）
    生产必做：`dynamic: strict` 或 `false` + 显式 mapping + index template。

14. **`object` 和 `nested` 的区别，为什么需要 nested？**
    普通 object 会在底层被「扁平化」成 `user.name: [a, b]`、`user.age: [1, 2]`，**丢失数组内元素的对应关系**，导致跨元素误匹配。
    `nested` 为每个子文档建独立的隐藏文档，保证查询时能锁在同一子文档内。代价是查询慢、聚合要 `nested` 声明；子文档多时可用 `flattened` 或把数据反规范化。

15. **parent-child（join 字段）和 nested 怎么选？**
    - nested：子文档跟父文档一起更新，查询快，适合「读多写少、子文档少」
    - join：父子是独立文档，可单独更新，但查询慢（要 join），一个索引只能一个 join 字段
    能用 nested 就别用 join。

16. **mapping 里 `index: false`、`doc_values: false`、`norms: false` 什么时候用？**
    - 只取不查的字段 → `index: false`
    - 不排序不聚合 → `doc_values: false`
    - 不关心长度归一化（如纯过滤字段）→ `norms: false`
    三者都能省大量磁盘。日志类场景普遍这么瘦身。

17. **索引该拆成多个小索引还是一个大索引？**
    - 按时间拆（`logs-2026.09`）：便于 ILM 滚动删除，冷热分层
    - 按业务域拆：mapping 不同、生命周期不同就拆
    - 拆太细：shard 过多，集群状态和查询扇出（fan-out）成本上升
    经验值：单 shard 10~50GB，shard 总数控制在节点数的 20 倍以内。

## 三、Query DSL

18. **`query context` 和 `filter context` 的区别？**
    - query：算相关性得分（`_score`），参与排序
    - filter：只判断「是/否」，**不算分、可缓存**（filter cache），性能好得多
    原则：能过滤的绝不查询。`bool.filter` 里的条件用于缩小候选集，`bool.must` 里只放真正影响相关性的条件。

19. **`term` 和 `match` 的区别，`term` 查 text 字段为什么查不到？**
    `term` 不分词，直接拿原词查倒排；`match` 会先分词再查。text 字段存的是分词后的 term，用 `term` 查 `"Hello World"` 整串必然查不到，应该查 `"hello"`（注意 analyzer 会做小写化）或改用 `match`。

20. **`bool` 的 must / should / filter / must_not 有什么讲究？**
    - `must`：必须满足，算分
    - `filter`：必须满足，不算分
    - `should`：满足加分。**如果 bool 里没有 must/filter，至少满足一个 should**
    - `must_not`：必须不满足，不算分（内部走 filter）
    - `minimum_should_match` 控制 should 命中几个
    顺序也影响性能：把过滤性最强的条件放前面没有实质作用，但**减少 must 的数量**是真能提分的。

21. **BM25 是什么，和 TF-IDF 差在哪？**
    BM25 = 词频饱和（term frequency 无限增长但增益递减）+ 文档长度归一化 + IDF。TF-IDF 的词频是线性累加的，长文档天然占便宜；BM25 用 `k1`、`b` 两个参数抑制这个偏差。ES 5.0 起默认 BM25。

22. **`boost` 和 `function_score` 的应用场景？**
    - `boost`：静态加权，比如 `title^3`
    - `function_score`：动态打分，比如「按销量、时间衰减、距离衰减」重排。电商搜索的排序策略基本靠它。

23. **`match_phrase` / `match_phrase_prefix` / `wildcard` / `regexp` 性能如何？**
    前两个相对可控（依赖位置信息），`wildcard` 和 `regexp` 是**性能杀手**，尤其是前导通配符 `*abc`，会退化成扫描所有 term。能用 `ngram` 或 `edge_ngram` 预处理就别用通配符。

24. **`terms` 查询里放几千个 ID 会怎样？**
    会撞 `indices.query.bool.max_clause_count`（默认 1024），报 `too_many_clauses`。解决方案：改用 `terms` lookup（从另一个索引读 ID 列表）、拆批查询、或改走 `ids` 查询 / 单独走 DB。

25. **聚合有哪几类？**
    - metric：`avg`/`sum`/`percentiles`/`cardinality`（基数，HyperLogLog++ 近似）
    - bucket：`terms`/`date_histogram`/`range`/`histogram`
    - pipeline：对聚合结果再聚合（`derivative`/`cumulative_sum`）
    坑：`terms` 聚合的 `size` 默认只返回 top 10，要拿全量得靠 `composite` 聚合分页。

26. **为什么 `cardinality` 聚合不准？**
    底层是 HyperLogLog++，用固定内存做基数估计，有 1~2% 误差。要精确值就加大 `precision_threshold`（代价是内存），或者换 `composite` + 遍历。

## 四、分页与深翻

27. **`from + size` 深分页为什么会 OOM？**
    每个 shard 都要取出 `from + size` 条并按 score 排序，协调节点再归并。翻到第 10000 页时，每个 shard 都要吐 10 万条，协调节点内存直接爆。默认 `index.max_result_window = 10000` 就是在拦这个。

28. **三种深翻方案怎么选？**
    | 方案 | 适用 | 缺点 |
    |---|---|---|
    | `from/size` | 前几页 | 深翻必炸 |
    | `scroll` | 离线全量导出 | 占资源、不能反映实时数据 |
    | `search_after` | **在线深翻首选** | 需要稳定的排序字段（如 `_id` + 时间） |
    - `search_after` 用法：先用一次查询拿到排序值，下一页带 `search_after: [last_sort_value]`，并且**必须用 PIT 或不变的排序**，否则数据变动会跳漏
    - 7.10 起推荐 **PIT（Point In Time）+ search_after**，替代 scroll 做实时深翻

29. **PIT 和 scroll 的核心差别？**
    PIT 保留一个时间点的索引视图，同时支持普通查询和 `search_after`，可用于实时场景；scroll 是一次性快照游标，适合大批量导出，会占用大量资源且不能并发修改。

## 五、写入与一致性

30. **`refresh_interval` 对写入性能影响多大？**
    默认 1s，写入时每 1s 生成一个新 segment，segment 越多 merge 压力越大。批量导入时先设 `-1` 并 `number_of_replicas: 0`，导完再改回来，吞吐能翻倍以上。

31. **`wait_for_active_shards`、`consistency`、`timeout` 是什么？**
    - `consistency`：写之前要求多少个 shard 可用的检查（1/quorum/all），7.x 起被 `wait_for_active_shards` 取代
    - `wait_for_active_shards=quorum`：保证写之前主分片有足够副本可用
    - 写请求超时不代表失败，可能只是响应慢，重试要靠幂等 ID

32. **写入一条文档的完整链路？**
    协调节点 → 按 routing 找到 primary shard → primary 写 translog + buffer → 同步给 replica → 等 replica 写完 → 返回。所以副本数越多，写延迟越高。

33. **怎么做到写入不丢？**
    `index.translog.durability=request`（每次请求 fsync translog）+ 至少 1 个副本 + `wait_for_active_shards=1` 以上。代价是每条写入都要 fsync，性能会掉。

34. **Bulk 的正确姿势？**
    - 单批 5~15MB，条数几千，不要一次几万条
    - 多线程并发（2~8 个客户端线程）比单线程大批更有效
    - **必须逐条检查响应里的 `errors`**，429（es_rejected_execution）要指数退避重试
    - `bulk` 里单条失败不会整体回滚

35. **`es_rejected_execution_exception`（429）怎么处理？**
    写队列（`thread_pool.write.queue_size`）满了。应对：客户端限流 + 退避重试、加大 bulk 间隔、减少并发、扩容节点。**千万别无脑重试**，会把集群压得更死。

## 六、Go 客户端

36. **Go 的 ES 客户端怎么选？**
    | 库 | 说明 |
    |---|---|
    | `github.com/elastic/go-elasticsearch/v8` | 官方，和 ES 8.x 对应，泛型化的 builder API |
    | `github.com/olivere/elastic` | 老牌，链式 API 最舒服，但 7.x 后基本停更 |
    | `github.com/elastic/go-elasticsearch/v7` | 老项目维护用 |
    新项目直接上官方客户端，模块版本和集群大版本对齐；别在同一个服务里混两个版本。

37. **官方 Go 客户端的基本用法？**
    ```go
    cfg := elasticsearch.Config{
        Addresses: []string{"http://localhost:9200"},
        Username:  "elastic",
        Password:  "changeme",
        Transport: &http.Transport{
            MaxIdleConnsPerHost:   100,
            ResponseHeaderTimeout: 5 * time.Second,
        },
    }
    es, err := elasticsearch.NewClient(cfg)
    _, err = es.Info()
    ```
    连多个节点用 `Addresses` 列表，客户端会自己做负载均衡和故障摘除。

38. **怎么优雅地写一个查询？**
    ```go
    var buf bytes.Buffer
    query := map[string]interface{}{
        "query": map[string]interface{}{
            "bool": map[string]interface{}{
                "filter": []interface{}{
                    map[string]interface{}{"term": map[string]interface{}{"status": "paid"}},
                    map[string]interface{}{"range": map[string]interface{}{
                        "created_at": map[string]interface{}{"gte": "now-7d"},
                    }},
                },
            },
        },
        "sort":  []interface{}{map[string]interface{}{"created_at": "desc"}},
        "size":  50,
    }
    if err := json.NewEncoder(&buf).Encode(query); err != nil {
        return err
    }
    res, err := es.Search(
        es.Search.WithContext(ctx),
        es.Search.WithIndex("orders-*"),
        es.Search.WithBody(&buf),
        es.Search.WithTrackTotalHits(true),
    )
    defer res.Body.Close()
    ```
    注意：`es.Search` 返回的 `res.IsError()` 为 true 时 body 是错误 JSON，**必须显式判错**，否则会把错误响应当数据解析。

39. **`TrackTotalHits` 是干什么的？**
    7.x 起 `hits.total` 默认只给「≥10000」这种近似值（`relation: gte`）以免全量计数。要精确值就 `WithTrackTotalHits(true)`，但大索引会明显变慢。

40. **怎么做 Go 的 bulk 写入？**
    ```go
    var buf bytes.Buffer
    for _, doc := range docs {
        meta := []byte(fmt.Sprintf(`{"index":{"_index":"orders","_id":"%s"}}`+"\n", doc.ID))
        data, _ := json.Marshal(doc)
        buf.Write(meta)
        buf.Write(data)
        buf.WriteByte('\n')
    }
    res, err := es.Bulk(bytes.NewReader(buf.Bytes()),
        es.Bulk.WithIndex("orders"),
        es.Bulk.WithRefresh("false"),
    )
    ```
    记得解析响应里的 `items[].index.error` 逐条确认。

41. **客户端常见错误类型怎么归类？**
    - 4xx：请求本身有问题（mapping 冲突、查询语法错），**重试无意义**
    - 429：队列满，可退避重试
    - 5xx：节点或分片问题，可重试
    - 网络超时：不确定是否写入成功，**靠文档 `_id` 幂等重试**

42. **Go 服务里怎么加超时和熔断？**
    ```go
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    res, err := es.Search(es.Search.WithContext(ctx), ...)
    ```
    每个请求都带 `WithContext`。再上层用 `gobreaker` 之类做熔断，ES 慢的时候别把业务线程池全占满。

43. **要不要自己封装一层 DAO？**
    要。把 index 名、query 构造、错误分类、重试策略、埋点（耗时、错误码、hits 数量）收在一层里。业务代码里散落的 `es.Search` 是后期最大的维护负担。

## 七、运维与调优

44. **JVM/内存怎么分？**
    堆内存别超过 32GB（开启压缩指针的阈值），一般是物理内存的一半且不超过 31GB。**剩下的一半留给 Page Cache**，Lucene 大量依赖文件缓存。堆设大了反而 GC 慢、缓存少。

45. **`indices.breaker` 熔断器是什么？**
    防止单个请求把内存吃光：`fielddata`、`request`、`in_flight_requests`、`accounting` 四类，超限直接报 `circuit_breaking_exception`。这是「聚合写太猛打挂集群」的兜底。

46. **shard 分配不出去（red/yellow）怎么排查？**
    ```bash
    GET _cluster/allocation/explain
    ```
    常见原因：磁盘水位（默认 85% 起不再分配、90% 尝试搬走）、节点刚重启还没恢复、副本数大于节点数。可临时 `PUT _cluster/settings {"transient":{"cluster.routing.allocation.enable":"all"}}`。

47. **滚动重启节点为什么要先关分片分配？**
    避免重启期间集群疯狂在节点间搬运分片（会产生巨量 IO）。流程：
    ```bash
    PUT _cluster/settings {"persistent":{"cluster.routing.allocation.enable":"primaries"}}
    POST _flush/synced          # 7.x 前；新版本用 POST _flush
    # 重启节点
    PUT _cluster/settings {"persistent":{"cluster.routing.allocation.enable":null}}
    ```

48. **零停机重建索引怎么做？**
    1. 建新索引（`orders_v2`）+ 新 mapping
    2. `POST _reindex?wait_for_completion=false` 从旧索引拷数据（`slices: auto` 并发）
    3. 校验文档数、抽样对比
    4. 用 alias 原子切换：`{"actions":[{"remove":{"index":"orders_v1","alias":"orders"}},{"add":{"index":"orders_v2","alias":"orders"}}]}`
    业务代码**永远只操作 alias**，这是前提。

49. **ILM（索引生命周期）解决什么问题？**
    hot → warm → cold → delete 的自动化：按时间/大小 rollover、迁到低配节点、force merge、最终删除。日志类场景标配。

50. **force merge 能提查询性能吗，代价是什么？**
    能，segment 少 → 需要打开的 Lucene reader 少 → 查询快、磁盘回收（清理已删文档）。但 force merge 是**极重 IO 操作**，只对**不再写入的索引**做，且只 merge 到 `max_num_segments=1`，别定期对手写索引做。

51. **节点角色怎么规划？**
    `master`、`data_content`/`data_hot`/`data_warm`/`data_cold`、`ingest`、`ml`、`transform`。生产至少 3 个专用 master、多个 data 节点，ingest 独立部署避免抢 CPU。单节点 Docker 起一个玩玩就全角色。

52. **搜索变慢怎么排查？**
    1. `GET _nodes/stats` 看 CPU / GC / thread_pool 的 rejected
    2. `GET orders/_search?profile=true` 看每个子查询耗时和执行计划
    3. `GET _tasks?detailed=true` 看有没有慢查询在跑
    4. 检查是否 `from+size` 深翻、是否 `wildcard`、是否 `terms` 爆 clause
    5. 检查 shard 是否过大、是否有 segment 过多

## 八、易错点速查

- 分片数创建后不可改，靠 alias + reindex 换索引
- `term` 查 text 字段查不到，是 analyzer 的锅
- 深翻别用 `from+size`，用 `search_after` + PIT
- `dynamic: true` 会导致 mapping explosion，生产用 template 显式声明
- bulk 要逐条检查 `errors`，429 要退避不要硬刚
- 堆内存不超过 31GB，剩下留给 Page Cache
- 业务只读 alias，直接写 index 名 = 以后没法零停机重建
- 批量导入前关副本 + 关 refresh，导入后恢复
- `hits.total` 7.x 起默认是近似值，别当精确数用
- 排序 / 聚合字段必须是 `keyword` 或 `numeric`，text 会报错或走 fielddata OOM
