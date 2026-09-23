```bash
spark-submit \
  --class "org.example.TestHiveSupport" \
  --master local[*] \
/mnt/c/Users/Janus_/Projects/work/maintenance_operation_procedure/hive_example/target/hive_example-1.0-SNAPSHOT.jar
```

```bas
spark-submit \
	--name "sync_task"   \
  --conf hive.metastore.uris=thrift://localhost:9083 \
  --master local[*] \
target/mpp-datasync-1.0.0.jar \
'{
"taskId":"12345",
"clickhouseTable":"janus.dwa_prd_pd_inst_day",
"hiveTable":"janus.dwa_prd_pd_inst_day",
"beachSize":"100",
"partitions":[
{"filed":"prov_id","value":"89"},
{"filed":"day_id","value":"20250430"}]}'
```

### 连接

```bash
```

### 运行流程

```bash
启动 Spark 应用程序

    使用 spark-submit 命令提交你的 JAR 包，这会触发整个 Spark 应用程序的启动。
    在这个过程中，Driver 程序开始运行，它负责初始化 SparkContext（对于 Spark 2.x 及以上版本，通常是 SparkSession），这是与集群交互的核心组件。

初始化 SparkContext/SparkSession

    这个步骤包括设置应用程序的各种配置参数、连接到集群管理器（例如 YARN）、以及准备执行环境。
    如果你指定了特定的资源要求（如 executor 内存大小、executor 核心数等），这些信息会被用来配置即将启动的 Executors。

连接到集群管理器

    Driver 通过集群管理器（比如 YARN）请求资源以启动 Executors。在这个例子中，YARN 的 ResourceManager 会分配容器（containers）来运行这些 Executors。

启动 Executors

    YARN 的 NodeManager 在集群中的各个节点上启动 Executor 实例。每个 Executor 是一个独立的 JVM 进程，负责执行实际的任务（tasks）。
    这些 Executors 启动后，它们会向 Driver 注册自己，表明准备好接受任务。

加载并执行 JAR 包中的代码

    当所有的 Executors 准备就绪后，Driver 开始执行你提交的 JAR 包中的主类（main class）。这个主类通常包含应用程序的入口点（即 public static void main(String[] args) 方法）。
    在这里，你定义的所有 RDD 操作、DataFrame/Dataset API 调用、SQL 查询等都会被执行，并形成一个或多个 DAGs（有向无环图）。

构建 RDD 血缘图

    Spark 根据你在 JAR 包中定义的操作链构建逻辑计划，即 RDD 血缘图或DAG。
    这些操作可以是转换操作（如 map, filter）或行动操作（如 collect, count）。

转换为物理执行计划

    DAGScheduler 将逻辑计划转化为物理执行计划，划分成不同的 Stage 和 Task。
    每个 Stage 包含一组具有相同依赖关系的 Tasks，这些 Tasks 可以并行执行。

调度 Tasks

    TaskScheduler 将 Tasks 分配给可用的 Executors 执行。考虑数据本地性原则，尽量让计算靠近数据所在的位置。

执行 Tasks

    Executors 接收到 Tasks 后开始执行。对于每个分区的数据，Executor 会执行相应的函数（如 map, reduce）。
    结果可能会被收集回 Driver 或者存储在外部系统中（如果指定了的话）。

结果返回

    对于行动操作（Action），如 collect() 或 saveAsTextFile()，结果会被发送回 Driver 程序进行进一步处理或输出。

清理资源

    当所有任务完成后，Driver 程序关闭 SparkContext，释放所有的 Executors 并清理相关资源。
    
    
    执行位置	描述
Driver 端	负责调度整个应用、创建 SparkContext、划分 DAG、提交任务等逻辑控制
Executor 端	负责执行实际的数据处理任务（如 map、filter、reduce 等）
```

### sql

```bash
spark-sql --conf spark.driver.memory=8G --conf spark.executor.memory=4G --num-executors 32
```

