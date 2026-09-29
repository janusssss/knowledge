# Java GC 面试常问问题（Go 为主 / Java 为辅 岗位）

> 定位：这类岗位面试官不会考得像 Java 后端那么深，但会拿 Go 当参照物问。重点在**分代模型 + 常见收集器 + 三色标记并发正确性 + 线上排查**，以及「为什么 Go 不分代而 Java 分代」「为什么 Java 敢移动对象而 Go 不敢」这几个跨界对照题。

---

## 一、内存结构与对象生死

1. **JVM 运行时内存分哪几块？哪些线程共享？**
   | 区域 | 共享性 | 存什么 | 会 OOM 吗 |
   |---|---|---|---|
   | 程序计数器 PC | 私有 | 当前字节码行号 | 不会（唯一不 OOM 的） |
   | 虚拟机栈 | 私有 | 栈帧、局部变量表、操作数栈 | StackOverflowError / OOM |
   | 本地方法栈 | 私有 | native 方法栈帧 | 同上 |
   | 堆 | **共享** | 对象实例、数组 | OOM: Java heap space |
   | 方法区（JDK8+ 元空间） | **共享** | 类元信息、运行时常量池、静态变量 | OOM: Metaspace |
   | 直接内存 | — | NIO DirectByteBuffer | OOM: Direct buffer memory |

   要点：JDK8 把永久代（PermGen）换成了**元空间**，类元信息挪到本地内存，因此「方法区」只是规范里的概念，实现是 Metaspace。静态变量（`static` 字段）和字符串常量池在 JDK7 就已经挪到堆里了。

2. **怎么判断一个对象该被回收？为什么 Java 不用引用计数？**
   用**可达性分析**：从 GC Roots 出发能走到的对象才存活。引用计数的死穴是**循环引用**——A 引用 B、B 引用 A，计数永不为 0，所以 Java 直接不用（Python 用引用计数 + 分代标记做补救，Go 干脆全用扫描）。

   GC Roots 包括：虚拟机栈中局部变量表引用的对象、方法区静态属性引用的对象、常量引用的对象、JNI 引用的对象、活跃线程、被 synchronized 持有的对象、JVM 内部引用（如基本类型对应的 Class 对象）。

3. **对象「死」一次就够了吗？**
   不是。可达性分析判定不可达后是**第一次标记**，随后判断是否需要执行 `finalize()`：没覆盖 `finalize()` 或已经执行过 → 直接回收；否则放进 F-Queue，由 Finalizer 线程执行（**不保证执行完**），执行中如果对象重新和 GC Roots 建立联系（自救），就能逃过这一轮。`finalize()` 已被 JDK9 标记废弃，实际没人依赖它。

4. **四种引用类型分别是什么场景？**
   - 强引用：`new` 出来的，只要有引用就不回收
   - 软引用 `SoftReference`：内存不足才回收 → 适合做缓存（现代实践更推荐 Caffeine 的 weakValues + 软引用已不流行）
   - 弱引用 `WeakReference`：下次 GC 必回收 → `WeakHashMap`、`ThreadLocalMap` 的 key
   - 虚引用 `PhantomReference`：拿不到对象，只用于回收时收到通知 → `DirectByteBuffer` 的 Cleaner 靠它释放堆外内存

   `ThreadLocal` 内存泄漏就出在这里：ThreadLocalMap 的 key 是弱引用（ThreadLocal 对象可以被回收），value 是**强引用**且挂在 Thread 上，线程不销毁则 value 一直活着。所以必须 `try/finally` 里 `remove()`。

5. **Java 的对象一定分配在堆上吗？**
   逻辑上都在堆，但 JIT 会做**逃逸分析**（`-XX:+DoEscapeAnalysis`，默认开）：
   - 对象不逃逸 → **标量替换**（拆成基本类型变量放栈上）、**栈上分配**、**锁消除**
   注意这几个只是优化手段、不改变语义，且可能因为分析失败而失效（比如方法大到不被 JIT 编译）。跟 Go 的**逃逸分析决定堆栈归属**不是一回事——Go 是编译器语义层面的，Java 只是运行时优化。

---

## 二、GC 算法与分代模型

6. **三种基础 GC 算法各自优劣？**
   | 算法 | 优点 | 缺点 | 用在哪 |
   |---|---|---|---|
   | 标记-清除 | 实现简单，不移动对象 | 碎片化、分配要靠空闲链表 | CMS |
   | 标记-复制 | 无碎片、分配是「指针碰撞」，极快 | 浪费空间、移动对象要改引用 | 新生代、Shenandoah/ZGC 的整理 |
   | 标记-整理 | 无碎片、不浪费空间 | 移动对象、停顿长 | 老年代（Serial Old、Parallel Old） |

7. **为什么分代？分代假说说了什么？**
   - 弱分代假说：**绝大多数对象朝生夕死**
   - 强分代假说：熬过越多次 GC 的对象越难死
   - 跨代引用假说：跨代引用只占极少

   于是新生代用复制（只扫存活的一小撮，成本与存活对象数量成正比），老年代用标记-整理/标记-清除（存活对象多，不适合复制）。**这是分代的根本理由：把「扫描成本」从「堆总大小」变成「存活对象量」。**

8. **新生代的默认比例？为什么是 8:1:1？**
   Eden : Survivor0 : Survivor1 = 8:1:1（`-XX:SurvivorRatio=8`），整个新生代占堆的 1/3（`-XX:NewRatio=2`）。
   理由：新生代存活率低，复制算法只需要一块**空的 Survivor** 作为 to-space，所以一开始就把 to-space 算成 10%，其余全给 Eden 用，空间利用率做到 90%，同时只要存活对象不超过 10% 就不会触发空间分配担保。

9. **Minor GC / Major GC / Full GC 都什么时候触发？**
   - **Minor / Young GC**：Eden 满
   - **Major / Old GC**：老年代满（语义上各家实现不一致，一般就等于 Full GC）
   - **Full GC**：老年代空间不足、元空间不足且无法扩容、`System.gc()`（除非 `-XX:+DisableExplicitGC`）、CMS 的 concurrent mode failure、G1 的 Humongous 分配失败、晋升时空间分配担保失败
   - **Mixed GC**：G1 特有，回收整个年轻代 + 部分收益高的老年代 Region

   Full GC ≈ 全堆 STW，是延迟杀手。线上排查「毛刺」第一件事就是确认有没有 Full GC。

---

## 三、对象分配与晋升

10. **一个对象从出生到进老年代的完整路径？**
    1. 优先在 **Eden** 分配
    2. 大对象直接进老年代（`-XX:PretenureSizeThreshold`，单位字节；**只对 Serial / ParNew 有效**，G1 用 Humongous 处理）
    3. 撑过一次 Minor GC 进 Survivor，年龄 +1（对象头 Mark Word 里 4 bit，最大 15）
    4. 年龄到 `-XX:MaxTenuringThreshold`（默认 15，CMS 默认 6）→ 晋升老年代
    5. **动态年龄判定**：Survivor 中相同年龄的对象总和 > Survivor 空间一半，则 ≥ 该年龄的对象直接晋升，不用等 15
    6. **空间分配担保**：Minor GC 前先看老年代最大连续可用空间是否 > **历次晋升对象的平均大小**，是才放心 Minor GC，否则改成 Full GC

11. **什么是 TLAB？为什么要有它？**
    Thread Local Allocation Buffer，每个线程在 Eden 里预先圈一块私有内存，分配对象时在自己的 TLAB 里指针碰撞，**不用加锁**。`-XX:+UseTLAB` 默认开启。TLAB 快满时线程再申请新的一块（浪费的一点点空间叫 refill waste）。

    Go 里对应的就是 **P 的 mcache + size class 分配器**，思路一模一样：线程（P）本地缓存、无锁分配。这一题非常适合主动接上 Go 对比。

12. **分配速率（Allocation Rate）为什么比堆大小更关键？**
    因为 Minor GC 的频率 ≈ 分配速率 / 新生代大小。堆调大只能降低频率，降不了晋升速率；**晋升速率**（每秒有多少字节从新生代活到老年代）才决定 Full GC 频率。压测时看 GC 日志里的 promotion rate，是调参最有价值的指标。

---

## 四、垃圾收集器逐个过

13. **有哪些收集器，各自怎么配对？**
    | 收集器 | 区域 | 算法 | 并发？ | 目标 |
    |---|---|---|---|---|
    | Serial | 新生代 | 复制 | 否 | 单核、客户端 |
    | Serial Old | 老年代 | 标记-整理 | 否 | 兜底、CMS 后备 |
    | ParNew | 新生代 | 复制 | 否（多线程 STW） | 配 CMS |
    | Parallel Scavenge | 新生代 | 复制 | 否（多线程） | **吞吐优先** |
    | Parallel Old | 老年代 | 标记-整理 | 否 | 配 Parallel Scavenge |
    | CMS | 老年代 | 标记-清除 | **是** | 低延迟（已移除） |
    | G1 | 整堆 | 复制+整理 | **是** | 可预测停顿（JDK9+ 默认） |
    | Shenandoah | 整堆 | 并发整理 | **是** | 低延迟，Brooks 指针 |
    | ZGC | 整堆 | 并发整理 | **是** | 超低延迟（<1ms） |
    | Epsilon | — | 不回收 | — | 压测/短命进程 |

    版本脉络：JDK8 默认 Parallel Scavenge + Parallel Old；JDK9 起默认 **G1**；JDK11 引入 ZGC（实验）；JDK12 Shenandoah；**JDK14 移除 CMS**；JDK15 ZGC/Shenandoah 转正；JDK21 引入**分代 ZGC**、虚拟线程。

14. **CMS 的四个阶段？**
    1. 初始标记（STW，只标 GC Roots 直接关联的对象，很快）
    2. 并发标记（与用户线程并行，最耗时但不 STW）
    3. 重新标记（STW，修正并发期间变动的引用，用**增量更新**）
    4. 并发清除（与用户线程并行）

    并发阶段用户线程还在跑，所以会持续产生新垃圾——叫**浮动垃圾**，只能等下次；同时要预留空间，`-XX:CMSInitiatingOccupancyFraction` 就是提前触发的阈值。

15. **CMS 的两个著名坑？**
    - **concurrent mode failure**：并发清除时老年代被填满，来不及了 → 退化成 Serial Old 做一次 Full GC（单线程、STW 巨长）。原因通常是晋升太快 / 碎片太多 / 阈值设置不当
    - **碎片化**：标记-清除不整理，`-XX:+UseCMSCompactAtFullCollection`（默认开）和 `-XX:CMSFullGCsBeforeCompaction` 只能在 Full GC 时补救

    还有：CMS 默认只能用 `-XX:ParallelGCThreads` 里约 1/4 的线程做并发标记，CPU 紧张时会「抢」用户线程。

16. **G1 的核心设计说清楚？**
    - 堆被切成约 2048 个 **Region**（1~32MB，`-XX:G1HeapRegionSize`），物理不分代、**逻辑分代**（Eden/Survivor/Old/Humongous，Region 角色动态变）
    - **Humongous**：单个对象 ≥ Region 一半大小，直接在连续 Region 分配，这类分配失败会触发 Full GC
    - **RSet（Remembered Set）+ Card Table**：每个 Region 维护「谁指向我」的卡片记录，反向扫描跨 Region 引用，避免扫全堆
    - **SATB（Snapshot-At-The-Beginning）+ 写前屏障**：并发标记期间保证不漏标
    - **Collection Set + 停顿预测模型**：`-XX:MaxGCPauseMillis`（默认 200ms）是目标不是承诺，G1 按 Region 回收收益排序，选一组能在目标时间内收完的
    - **Mixed GC** 由 `-XX:InitiatingHeapOccupancyPercent`（默认 45%）触发；JDK9+ 默认开启自适应 IHOP
    - Young GC 是 **STW 且有复制**，所以 G1 不是「无停顿」

17. **ZGC 为什么能到亚毫秒停顿？**
    - **染色指针（Colored Pointers）**：把标记位（Marked0/Marked1/Remapped/Finalizable）直接塞进 64 位指针的高位，标记结果存在引用上而不是对象头，**不需要访问对象**就能判断状态
    - **读屏障**：应用线程读引用时顺便修正指针（自愈），把标记/整理工作分摊到用户线程
    - **并发整理**：对象移动时旧地址留下转发指针，读屏障负责纠正
    - 代价：读屏障拖慢应用吞吐、内存占用更高（指针多态 + 不支持压缩指针 `-XX:-UseCompressedOops`）、依赖 64 位
    - JDK21 的**分代 ZGC** 补上了「不分代」这个短板（不分代时每次都得扫全堆，分配速率一高就撑不住）

    Shenandoah 思路相近，但用 **Brooks 转发指针**（对象头多一个字段指向自己），不用染色指针，因此不要求特定平台。

---

## 五、三色标记与并发正确性（重点)

18. **三色标记的三色分别代表什么？**
    - 白色：未访问（标记结束仍是白 = 不可达 = 可回收）
    - 灰色：自己已被访问，但引用的对象还没扫完（在标记栈里）
    - 黑色：自己和所有引用都已经扫完（不用再扫）

19. **并发标记为什么会「漏标」？漏标的充要条件是什么？**
    两个条件同时成立才会漏标：
    1. 黑色对象**新增**了指向白色对象的引用（写屏障没能拦住）
    2. 灰色对象**删掉**了指向该白色对象的引用（这条链断了，以后没人再去扫它）

    Go 1.8 之前只有插入写屏障（Dijkstra），所以要 STW 重扫栈；Java 的 CMS 用增量更新解决条件 1，G1 用 SATB 解决条件 2。

20. **增量更新（Incremental Update）和原始快照（SATB）的区别？**
    | | 增量更新 | 原始快照 SATB |
    |---|---|---|
    | 谁用 | CMS | G1、Shenandoah |
    | 记录什么 | 并发中**新增**的「黑→白」引用 | 并发中**被删除**的旧引用 |
    | 重新标记时 | 以这些黑色对象为根**重新扫描** | 把这些旧引用当作活着的（快照） |
    | 语义 | 保证不漏标当前活对象 | 保证不漏标**标记开始那一刻**的活对象，代价是有浮动垃圾 |
    | 屏障类型 | 写后屏障（post-write barrier） | 写前屏障（pre-write barrier）+ 写后 |

    G1 的 SATB 队列是**每线程一个**，满了就全局并发队列，溢出了还得 STW 处理，这是 G1 停顿的一个隐藏来源。

21. **Go 的混合写屏障解决了什么问题？**
    Go 1.7 之前是「Dijkstra 插入屏障 + 标记结束时 STW 重扫栈」，因为栈上对象不能加写屏障（开销太大），必须 STW 修正。
    Go 1.8 引入**混合写屏障**：GC 开始时把栈上对象**全标黑**（栈在本次 GC 期间新引用的对象也标灰），此后栈不需要重扫，于是 **STW 只剩 mark setup 和 mark termination 两个极短的阶段**，不再有「重扫栈」这种长停顿。

    这一题如果被问到，能把 Go 1.8 的改动和 Java 的粘性写屏障/增量更新对比讲，基本就稳了。

---

## 六、GC 日志、参数与线上排查

22. **怎么开 GC 日志？**
    ```bash
    # JDK8
    -XX:+PrintGCDetails -XX:+PrintGCDateStamps -Xloggc:/var/log/app/gc.log
    # JDK9+ 统一日志
    -Xlog:gc*,gc+heap=debug,gc+age=trace:file=/var/log/app/gc.log:time,uptime,level,tags:filecount=10,filesize=64m
    ```
    JDK9 之后 `PrintGCDetails` 这些参数全被移除，只剩 `-Xlog`。生产一定要带 `filecount`/`filesize` 轮转，不然日志把磁盘写满。

23. **GC 日志主要看什么？**
    - **频率**：Young GC 多久一次、Full GC 多久一次
    - **单次停顿**：`[Times: user=0.02 sys=0.00, real=0.01 secs]` 里的 real 才是墙钟
    - **晋升速率**：每次 GC 后老年代涨了多少
    - **GC 后老年代是否回落**：Full GC 后老年代几乎不降 → 大概率**内存泄漏**；降了但很快又满 → 容量不足或阈值不合理
    - **Humongous 分配**：G1 下频繁出现 Humongous allocation 说明有大对象

24. **常用参数背几个？**
    ```bash
    -Xms4g -Xmx4g                    # 设成相等，避免动态扩容带来的抖动
    -XX:MetaspaceSize=256m -XX:MaxMetaspaceSize=256m
    -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/data/dump
    -XX:+UseG1GC -XX:MaxGCPauseMillis=200
    -XX:ParallelGCThreads=8 -XX:ConcGCThreads=2
    -XX:+DisableExplicitGC           # 防 RMI 之类的显式 System.gc()
    ```
    调参顺序：先确认没有泄漏 → 再看停顿目标 → 再看吞吐，别一上来就调 `-XX:NewRatio`。

25. **有哪几种 OOM？分别怎么定位？**
    | 报错 | 原因 | 手段 |
    |---|---|---|
    | `Java heap space` | 堆不够 / 泄漏 / 大对象 | `jmap -dump` + MAT 找 dominator |
    | `GC overhead limit exceeded` | GC 占了 98% 时间却回收不到 2% 的堆 | 同堆 OOM，本质是堆太小或泄漏 |
    | `Metaspace` | 类加载太多（动态代理、CGLib、热部署） | 看 `jstat -gc` 的 MC/MU，`-XX:MaxMetaspaceSize` |
    | `Direct buffer memory` | 堆外内存没释放 | `-XX:MaxDirectMemorySize`，看 NIO/Netty |
    | `unable to create new native thread` | 线程数到系统上限 / 栈太大 | 数线程、降 `-Xss`、查 ulimit |
    | `Requested array size exceeds VM limit` | 数组长度超 int 上限附近 | 逻辑问题 |

    注意 `StackOverflowError` 是栈溢出，和 OOM 不是一回事，但深递归也耗内存。

26. **线上 CPU 100% / 频繁 Full GC 的排查路线？**
    ```
    1. jstat -gcutil <pid> 1000      # 先看是不是 GC 问题，FGCT 涨得快就是
    2. top -Hp <pid>                 # 揪出占 CPU 的线程
    3. printf '%x\n' <tid>           # 转十六进制
    4. jstack <pid> | grep -A 30 <hex>  # 看线程栈，找业务热点
    5. jmap -histo:live <pid> | head -20  # 看谁在堆里暴涨
    6. jmap -dump:live,format=b,file=/tmp/heap.hprof <pid>
    7. MAT / arthas heapdump 分析支配树与 GC Roots 引用链
    ```
    arthas 一把梭：`dashboard`（看 GC 与线程）、`thread -n 3`（最忙的 3 个线程）、`trace`/`watch` 定位慢方法。

27. **JDK21 虚拟线程对 GC 有什么影响？**
    虚拟线程的**栈不放在 OS 栈上，而是以 chunk 形式存在 Java 堆里**，可动态挂载/卸载。所以百万虚拟线程 = 堆里多了一堆栈 chunk，堆压力会明显上升（尤其是 `-XX:MaxGCPauseMillis` 敏感的场景）。好处是不再需要「线程池 + 异步回调」那套，坏处是 ThreadLocal 要谨慎（改用 ScopedValue）。

    Go 的 Goroutine 栈也是堆上的、按需复制扩容，本质上是一模一样的设计——**这一对照很值得主动说**，会被认为真的理解 go runtime。

---

## 七、Go GC 对照（这类岗位最可能问的跨界题）

28. **Go 的 GC 大概怎么工作？**
    - **并发三色标记 + 混合写屏障**，非分代（Go 1.19 一直在讨论，但直到现在主线仍不分代）
    - **不移动对象（non-moving）**：不做压缩整理，靠 size class + span 缓解碎片
    - **STW 只有两处**：mark setup（开启写屏障）和 mark termination（关闭写屏障），通常几百微秒
    - 触发条件：堆大小达到 `live heap * (1 + GOGC/100)`，`GOGC` 默认 100；Go 1.19 起支持 `GOMEMLIMIT` 软内存上限，两者共同作用
    - 辅助标记（mark assist）：分配太快的 goroutine 会被抓去帮忙标记，**用延迟换内存**
    - `GODEBUG=gctrace=1` 看每次 GC 的 `scan/heap/mark/sweep` 时间

29. **为什么 Go 不做分代，Java 做？**
    - 分代的前提是能**低成本识别并追踪代际引用**（Java 有 Card Table + 写屏障，Go 的写屏障成本本来就高）
    - 分代 + 复制意味着**移动对象**，移动就要修正所有引用，而 Go 的栈是**可增长的**、指针可能在栈上、还有 unsafe.Pointer / cgo 边界，精确移动的工程代价极高
    - Go 的目标是「**低延迟 + 简单**」，不是「高吞吐」，所以宁可用「不分代但全程并发」这条路，把延迟压到亚毫秒级，代价是吞吐和内存放大比 Java 高
    - Go 团队也承认：分代如果要做得小停顿时，收益不足以抵消复杂度（Go 1.19 起大量工作是分层 GC 之外的「GC 开销随栈数量线性增长」问题）

30. **为什么 Java 敢移动对象，Go 不敢？**
    Java 能移动，因为：
    - 栈是**固定大小的、有精确栈图的（precise stack map，栈上哪一格是指针编译器完全知道）**，需要全 STW（或安全点）来统一修正栈上引用
    - 对象引用不暴露给外部（不许裸指针运算），JNI 通过句柄/pin 处理
    - 因此可以复制、可以整理，碎片问题天然解决

    Go 不移动，因为：
    - 栈是可增长的（栈扩容时会复制栈并修正指针），跟 GC 移动叠在一起会非常复杂
    - 存在 `unsafe.Pointer`、`cgo` 传指针、`runtime.KeepAlive` 等漏网之鱼
    - 结论：**用「不整理」换「实现简单 + STW 短」**，碎片靠 size class 分配器吸收

31. **`GOGC` / `GOMEMLIMIT` vs `-Xmx` / `-Xmn` 的对照？**
    | | Go | Java |
    |---|---|---|
    | 触发依据 | 存活堆 × (1+GOGC) 或内存上限 | 各个代是否满（Eden/老年代） |
    | 主要旋钮 | `GOGC`、`GOMEMLIMIT`、`GOMAXPROCS`（影响标记并发度） | `-Xmx/-Xms/-Xmn/-XX:MaxGCPauseMillis/-XX:SurvivorRatio` |
    | 调参空间 | 很小，基本两个 | 很大，能针对延迟/吞吐分场景 |
    | 内存下限 | `GOGC=off` + `GOMEMLIMIT` 就是「只在触顶时 GC」 | 没这么直接的开关 |

    Go 的 `GOMEMLIMIT` 相当于「软版 `-Xmx`」：不硬性阻止超限，但接近时 GC 会变积极防止被 OOM Killer 干掉。Java 的 `-Xmx` 是硬上限，超了直接 `OutOfMemoryError`。

32. **弱引用 / 终结器的对照？**
    - Java：四种引用类型齐备，`Cleaner`（JDK9 起）替代 `finalize`，`DirectByteBuffer` 就靠它释放堆外内存
    - Go：**没有弱引用**，只有 `runtime.SetFinalizer`，而且不保证被执行、还可能让对象多活一轮甚至永久存活
    - 这题如果问「Go 里怎么实现缓存过期」——只能靠外部 LRU（`container/list` 或 `groupcache/lru`）、time.After 的定时清理或 `sync.Map` + 时间戳，没有语言级弱引用兜底

33. **「你们 Go 服务的内存曲线和 Java 服务有什么不同？」**
    好的回答方向：
    - Go 的内存是**锯齿状但底噪高**：GC 结束不会把内存还给 OS（除非 scavenger 回收），RSS 常驻在高水位；Java 也类似（堆提交后不轻易归还）
    - Go 因为没有整理，长时间运行后碎片 + 高水位更明显，容器里要按**峰值 + 30%** 留 `limits`
    - Java 因为分代，短命对象不参与老年代 GC，所以「分配速率高但存活少」的场景 GC 停顿反而可控
    - 真正决定停顿的是 **live heap 大小**，不是总堆大小——两边都一样

---

## 八、答题策略（Go 为主 Java 为辅）

34. **面试官只会有几个「真考点」**
    大概率是：内存分区 → 分代与晋升 → G1/CMS 原理与区别 → 三色标记并发正确性 → 一次线上问题怎么查。ZGC 细节、Shenandoah 的 Brooks 指针这种属于「说得出加分，说不出不扣分」。

35. **主动对比 Go，但要准确**
    - 可以说：Go 的 mcache/TLAB 都是线程本地无锁分配；Go 1.8 混合写屏障对应 Java 的增量更新/SATB；Goroutine 栈与虚拟线程栈都在堆里按需增长
    - 别硬套：Go 的「不分代」不是技术做不到，是**权衡**；Java 的分代是「用写屏障 + 精确栈图换吞吐」
    - 别说的：别把 Go 的 GC 说成「标记-清除」（Go 也扫 span，不整理不等于不管理）；别说 Go 有 STW 也很长（通常 sub-ms）

36. **被问到「你们为什么选 Go 不选 Java」这种开放题**
    落在这几条上，面试官最爱听：
    - 部署：静态二进制 vs JVM 启动/预热（JIT 冷启动、`-Xmx` 与容器配额）
    - 延迟：Go 的 sub-ms STW 对比 Java 的 G1 200ms 目标（但 Java 有 ZGC 扳回一局）
    - 生态：Java 的 Kafka/ES 客户端、Netty、Spring 生态成熟；Go 的并发模型和内存占用更省
    - 团队：运维复杂度、招聘成本、现有技术栈
    别忘了承认 Java 的优势，一句「Java 在 GC 调优空间和成熟中间件上确实强，Go 是拿调优自由度换简单」收尾最安全。

37. **反问环节可以问什么**
    - 线上 JVM 是哪个版本、默认收集器是什么（JDK8 vs 11/17/21 结论完全不同）
    - 有没有 CDS/AppCDS、JIT 预热策略
    - 采集链路是 OpenTelemetry 还是 Prometheus JMX Exporter
    - 数据库/中间件那些 Java 组件的 GC 是谁在管

---

## 九、GC 的完整流程（从触发到结束）

38. **先把通用主干记住：任何一次 GC 都走这九步**
    ```
    ① 触发判定      堆/代的空间水位、晋升失败、显式调用
    ② 选定回收范围  新生代 / CSet / 整堆
    ③ 到达安全点    所有线程停在 Safepoint（这段叫 TTSP，常被忽略）
    ④ 根枚举        从 GC Roots 出发，扫描栈、寄存器、静态区、JNI
    ⑤ 标记          三色标记，并发时靠读写屏障保证正确性
    ⑥ 引用处理      强→软→弱→虚，死亡引用入 ReferenceQueue
    ⑦ 回收动作      复制（新生代）/ 清除（CMS）/ 整理（Old、G1 的 CSet）
    ⑧ 元数据更新    对象年龄、RSet、Card Table、空闲链表、Region 状态、TLAB
    ⑨ 退出安全点    恢复用户线程，继续分配
    ```
    被问「讲一下 GC 流程」，先给这九步，再按「你问的是哪种 GC」展开，比直接背 CMS 四阶段稳得多。

39. **一次 Young GC（Parallel Scavenge / ParNew）的完整流程？**
    1. **触发**：Eden 分配不下新对象（且 TLAB 也 refill 不出来）
    2. **担保检查**：老年代最大连续可用空间 > 历次晋升平均大小？不满足就直接 Full GC
    3. **STW**：所有用户线程停在安全点
    4. **根枚举**：扫描 GC Roots，重点是各线程栈上的局部变量和寄存器（这是 STW 必须存在的原因——栈上的指针不能并发被改）
    5. **找活对象**：从根出发，沿对象引用遍历（新生代用 Card Table 处理老年代 → 新生代的跨代引用，只需扫描 dirty card）
    6. **复制/晋升**：把存活对象复制到空的 Survivor（to-space）；放不下或年龄到阈值就复制到老年代；记住所有引用都要改到新地址
    7. **清空**：Eden 和 from-space 整体清空（不需要一个个标，复制算法里非存活自然被丢弃）
    8. **更新年龄**：Survivor 中对象 Mark Word 里的分代年龄 +1，同时做**动态年龄判定**
    9. **恢复**：新 TLAB 重新划分，用户线程继续跑

    关键理解：**新生代的回收成本 ∝ 存活对象数量，而不是堆大小。** 所以「Eden 调大只会让 GC 频次降低，不会让它变慢」。

40. **一次 G1 Young GC 的完整流程？和 ParNew 差在哪？**
    1. **触发**：Eden 的 Region 用完
    2. **选 CSet**：Young GC 的 CSet 只包含全部 Eden + Survivor Region（`-XX:G1NewSizePercent` / `G1MaxNewSizePercent` 约束年轻代占比 5%~60%）
    3. **STW + 根枚举**：多了两件事——扫 **RSet**（谁指向我）和 SATB 队列
    4. **复制/晋升**：存活对象复制到新的 Survivor Region 或 Old Region，**跨 Region 的引用靠 RSet 精确修正**，同时更新 Card Table
    5. **清空 CSet**：Region 被整体回收，收回的空 Region 放回空闲列表（free list）
    6. **可能要顺带做 Initial Mark**：如果 IHOP（默认 45%）到了，这次 Young GC 会挂上并发标记的开始阶段

    与 ParNew 的本质差异：G1 的回收单位是 **Region** 而不是整个代，所以它每次都能算出「收哪些 Region 收益最高」，这就是停顿可预测的来源。

41. **G1 一次完整的并发标记周期（concurrent marking cycle）？**
    | 阶段 | STW？ | 干什么 |
    |---|---|---|
    | Initial Mark | **是**（搭在 Young GC 上） | 标 GC Roots 直接关联的对象，顺便开启 SATB 写屏障 |
    | Root Region Scan | 否 | 并发扫描 Survivor Region 里指向老年代的引用（这段必须在下一次 Young GC 前做完） |
    | Concurrent Mark | 否 | 并发遍历整个堆，用 SATB 快照保证不漏标，标记栈溢出就退化重扫 |
    | Remark | **是** | 处理 SATB 队列残留、引用处理，最耗时也最不可控的一步 |
    | Cleanup | **是**（很短） | 统计各 Region 存活率、回收「全是垃圾」的 Region、决定下次要不要 Mixed GC |

    结束后进入 **Mixed GC** 阶段：回收全部年轻代 + 部分（收益高的）老年代 Region，混合回收的轮次受 `-XX:G1MixedGCCountTarget` 和 `G1HeapWastePercent` 控制。

    所以完整时间线是：**Young GC × N → 并发标记 → Young GC → Mixed GC × M → 再回 Young GC**。这也解释了为什么 G1 的「一次 Full GC」往往是 Mixed GC 来不及收干净导致的兜底。

42. **CMS 的完整时间线（带 STW 标记）？**
    ```
    ┌─初始标记(STW,极短)─┐
    │                    └──并发标记──────────────────────┐
    │                                                      ├─重新标记(STW)─┐
    │                                                      │               └──并发清除────────→
    │                                                      │
    用户线程 ──────────────────────────────────────────────────────────────────────────→
              ↑ 老年代达到 CMSInitiatingOccupancyFraction 时触发
    ```
    - 初始标记：只标 GC Roots 的直接关联对象（快，因为不递归往下走）
    - 并发标记：从这些根出发遍历整个可达图，最耗时的一个阶段，但用户线程在跑
    - 重新标记：**修正并发期间的引用变动**（增量更新记录的黑→白引用），是 CMS 最长的 STW 阶段
    - 并发清除：边跑边清，产生**浮动垃圾**；如果此时老年代被填满 → **concurrent mode failure**，退化 Serial Old 做 Full GC

43. **一次 Full GC 的完整流程？为什么它那么慢？**
    1. 所有线程 STW（往往是全部线程，不只应用线程）
    2. 根枚举（含元空间、所有线程栈）
    3. 标记整堆可达对象（老年代存活率高，标记代价大）
    4. 整理或清除（若整理，还要移动对象 + 修正所有引用，成本 ∝ 存活对象数）
    5. 收尾：引用处理、类卸载（G1 在 Full GC 时才会尝试卸载无用类）、元空间/永久代清理
    6. 恢复

    慢的根本原因：**要扫的活对象多 + 全程不能并发 + 可能还要移动**。所以「减少 Full GC」的正确姿势是减少晋升（别让短命对象活得太久），而不是一味堆大堆。

44. **STW 到底怎么实现的？安全点和安全区域是什么？**
    - JVM 要 STW 时不是「去暂停线程」，而是**设置一个标志位**，各线程跑到**安全点（Safepoint）** 时主动检查并挂起 → 叫「主动式中断」
    - 安全点位置由 JIT 插桩决定：方法调用、循环回跳（back edge）、异常跳转、native 方法返回
    - 已经在阻塞状态（sleep、wait、IO）的线程，处在**安全区域（Safe Region）**，可以直接放行，不用等它跑回安全点
    - **注意**：GC 日志里的 `real=0.05 secs` 是 GC 本身耗时，日志里 `Total time for which application threads were stopped` 或安全点日志里的时间是 **TTSP（Time To Safepoint）**，两者相加才是用户真正感知的停顿。长 `while` 循环、大数组拷贝、JNI 里的死循环都会让 TTSP 爆表——**「GC 明明很快但接口还是卡」十有八九是 TTSP**

45. **GC 什么时候处理 Reference？**
    标记结束后、回收动作前，按 **强 → 软 → 弱 → 虚** 的顺序处理，并把死亡引用入队到 `ReferenceQueue`：
    - 软引用只在「内存不足」时清除（`-XX:SoftRefLRUPolicyMSPerMB` 控制）
    - 弱引用只要标记时不可达就被清
    - 虚引用 + `Cleaner` 负责堆外内存释放（`DirectByteBuffer`）
    - 处理阶段可并行：`-XX:+ParallelRefProcEnabled`，**Netty/大量 WeakReference 的服务必开**，能省掉一大截 STW

46. **对象年龄与 RSet 是在流程的哪一步更新的？**
    - 年龄：复制到 Survivor 时 +1（Mark Word 里 4 bit，所以最大 15）
    - RSet / Card Table：**对象移动或引用被改写时**更新，写屏障负责记录脏卡；GC 期间修正引用后也要同步更新 RSet
    - TLAB：GC 结束后线程重新申请 TLAB（这一步也可能触发分配失败）

47. **对照：Go 的一次 GC 完整流程？**
    ```
    GC start
      → mark setup（STW，开启混合写屏障，很短）
      → concurrent mark（25% CPU 起，含 mark assist 辅助标记）
      → mark termination（STW，关闭屏障、统计）
      → concurrent sweep（懒惰清扫，分配时才清 span）
      → GC end（下次触发点 = live heap × (1+GOGC/100) 或 GOMEMLIMIT）
    ```
    - 没有「分代」「晋升」「复制」这些步骤，只有一轮扫全堆
    - **没有独立的引用处理阶段**（Go 没有弱引用，`SetFinalizer` 的对象要等下一轮）
    - sweep 是**惰性**的，跟 Java 的「清空」是两种思路
    - Go 1.19+ 有 `GOMEMLIMIT` 作为软上限，接近时提前开 GC

    一句话对照：**Java 的流程是「分多轮、每轮只碰一部分」，Go 的流程是「一轮扫全部、但全程并发」**。

48. **面试被问「讲一下 GC 流程」怎么组织答案？**
    三步走：
    1. 先给主干九步（触发 → 安全点 → 根枚举 → 标记 → 引用处理 → 回收 → 重置）
    2. 再分场景：年轻代（复制、晋升）、老年代（CMS 四阶段 / G1 Mixed）、Full GC（整堆 STW）
    3. 最后补一句容易漏的：**TTSP、引用处理、RSet 更新**——这三个点说得出，面试官会认为你真看过日志

    如果只有一句话的位置：**「判活靠可达性分析，收垃圾看分代，并发靠屏障，暂停看安全点。」**

---

## 附：一分钟速记

- 内存分区：PC / 栈 / 本地栈 私有，堆 / 元空间 共享
- 存活判定：可达性分析 + GC Roots；不用引用计数（循环引用）
- 算法：复制（新生代）、标记-清除（CMS）、标记-整理（老年代）
- 分配：TLAB → Eden → Survivor（8:1:1）→ 年龄 15 或动态年龄 → 老年代；大对象直接进老年代
- 收集器：JDK8 默认 Parallel，JDK9+ 默认 G1，JDK14 移除 CMS，JDK21 分代 ZGC
- CMS 四阶段：初标 STW / 并发标记 / 重标 STW / 并发清除；坑 = 浮动垃圾 + 碎片 + CMF
- G1：Region + RSet + SATB + 停顿预测 + Mixed GC；Humongous ≥ Region/2
- 并发正确性：漏标 = 黑→白新增 + 灰→白删除；增量更新（CMS）vs 原始快照（G1）
- 排查：jstat -gcutil → top -Hp → jstack → jmap -histo/-dump → MAT/arthas
- Go 对照：并发三色 + 混合写屏障、非分代、不移动、STW 只在 setup/termination、GOGC + GOMEMLIMIT
- GC 流程：触发 → 选范围 → 安全点(STW) → 根枚举 → 三色标记 → 引用处理 → 复制/清除/整理 → 更新年龄与 RSet → 退出
- 停顿 = GC 自身时间 + **TTSP**（安全点等待），日志里 real 只是前者
