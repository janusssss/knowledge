# maven

### 本地更新maven仓库

```bash
mvn install:install-file -Dfile=ojdbc6.jar -DgroupId=com.oracle -DartifactId=ojdbc6 -Dversion=11.2.0.3 -Dpackaging=jar -DgeneratePom=true
```

### 拉取依赖

```bash
 mvn dependency:resolve
```

### Lifecycle

```bash
clean: 清理目标目录。
validate: 确认项目配置正确无误。
initialize: 初始化构建环境。
generate-sources 和 process-sources: 生成并处理源代码。
generate-resources 和 process-resources: 生成并处理资源文件。
compile: 编译源代码。
test-compile: 编译测试源代码（如果有）。
process-test-resources: 处理测试资源文件。
test: 运行单元测试并生成报告。
package: 打包项目。
install: 安装打包后的项目到本地仓库。
```



## spring cloud 启动过程

pring Cloud作为构建于Spring Boot之上的框架，用于简化分布式系统的开发。其启动流程与Spring Boot非常相似，但包含了额外的步骤来初始化和配置微服务架构中的组件，如服务发现、配置管理、断路器等。以下是Spring Cloud应用的一般启动流程：
1. 初始化Spring Boot应用程序

    加载主类：Spring Boot应用通常有一个带有@SpringBootApplication注解的主类（通常是public static void main(String[] args)方法所在的类）。这个注解包含了@Configuration, @EnableAutoConfiguration, 和 @ComponentScan注解。

    自动配置：基于classpath下的依赖，Spring Boot会自动配置你的应用程序。例如，如果HSQLDB在classpath中，而你没有手动配置任何数据库连接，Spring Boot会自动配置一个内存数据库。

2. 加载外部化配置

    读取配置文件：Spring Boot首先从bootstrap.yml或bootstrap.properties加载配置，这对于使用Spring Cloud Config Server的应用来说尤为重要，因为这些配置文件用于指定Config Server的位置和其他初始配置信息。

    配置属性替换：如果使用了占位符（如${...}），此时会根据实际配置进行值的替换。

3. Spring Cloud组件初始化

    服务发现：如果你使用了Eureka或其他服务发现机制，Spring Cloud会在这一阶段尝试注册自身到服务发现服务器，并查找其他依赖的服务。

    配置管理：如果配置了Spring Cloud Config Server，Spring Boot将尝试从Config Server获取应用配置。这一步骤通常发生在应用程序的早期阶段，甚至在Spring上下文完全加载之前。

    断路器和负载均衡：初始化Ribbon（客户端负载均衡器）和Hystrix（断路器库）等组件，以便后续的请求可以利用这些功能。

4. 应用程序上下文刷新

    Bean定义加载：扫描并加载所有被@Component, @Service, @Repository, @Controller等注解标记的类作为Spring Beans。

    执行Bean工厂后处理器：允许修改Bean定义，包括添加、移除或修改Bean定义。

    实例化Beans：创建所有非懒加载的单例Beans，并解决它们之间的依赖关系。

5. 启动完成

    调用CommandLineRunner和ApplicationRunner接口的实现类：如果有的话，Spring Boot会调用这些类的run方法，允许开发者在Spring Application完全启动之后执行一些操作。

    健康检查和监控：Spring Boot Actuator提供了一系列端点来监控和管理应用，这些端点现在也已可用。

# javac

```bash
javac -d classes/ src/*.java
```

```bash
javac [options] [source files]
# [options]：可选参数，指定编译选项
-d 选项可以指定编译后的 .class 文件存放的目录
-verbose 选项可以看到更详细的编译过程信息
-cp 或 -classpath 来指定包含所需类文件或 JAR 文件的路径
# [source files]：要编译的一个或多个 Java 源代码文件
```

### 方法分类

| 方法类别         | 翻译     | 概述                                                         |
| ---------------- | -------- | ------------------------------------------------------------ |
| Instance Methods | 实例方法 | 定义在类中的普通方法，需要通过该类的一个实例来调用           |
| Abstract Methods | 抽象方法 | 没有具体实现的方法，只有方法签名，通常用于声明某个行为而没有提供具体的实现 |
| Default Methods  | 默认方法 | 接口中为方法提供一个默认实现                                 |



### 通配符

| 符号            | 作用                                           |
| --------------- | ---------------------------------------------- |
| `<?>`           | 表示这个类型存在，但我们不关心它是什么具体类型 |
| `<? extends T>` | 某种 T 的子类类型                              |
| `<? super T>`   | 某种 T 的父类类型                              |

### 函数接口

| 接口                                | 作用             |
| ----------------------------------- | ---------------- |
| Runnable                            | 无参无返回值     |
| Function<String, Integer>           | 单个入参有返回值 |
| Consumer<String>                    | 有入参无返回值   |
| Supplier<String>                    | 无入参有返回值   |
| BiFunction<String, String, Integer> | 两个入参有返回值 |
|                                     |                  |

```java
// 自定义方法接口
@FunctionalInterface
public interface TriFunction<T, U, V, R> {
    R apply(T t, U u, V v);
}
```

---


# Java GC 概述

## 什么是 GC
GC（Garbage Collection，垃圾回收）是 JVM 自动管理堆内存的机制，负责回收不再被引用的对象，释放内存空间。

## JVM 堆内存分区
JVM 采用**分代收集**思想，把堆分成不同区域，因为大多数对象是"朝生夕灭"的：

- **年轻代（Young）**：Eden + 两个 Survivor（S0/S1）。存放新对象，GC 频繁但快。
- **老年代（Old）**：存放存活多次仍被引用的对象，GC 少但耗时长。
- **元空间（Metaspace，JDK8+）**：存放类元数据，不在堆内（取代了永久代）。

## 回收算法
1. **标记-清除（Mark-Sweep）**：标记存活对象，清除其余。缺点：产生内存碎片。
2. **复制（Copying）**：把存活对象复制到另一块区域，用于年轻代（Eden 存活对象复制到 S0/S1）。无碎片，但浪费一半空间。
3. **标记-整理（Mark-Compact）**：标记后把存活对象向一端移动。用于老年代，无碎片。

## Major Collector 演进
| GC | 特点 |
|---|---|
| **Serial** | 单线程，暂停时 Stop-The-World，适合单核/小内存 |
| **Parallel / Parallel Old** | 多线程并行收集，追求吞吐量，JDK8 默认 |
| **CMS** | 并发标记清除，低停顿，但产生碎片、并发失败触发 Full GC |
| **G1（JDK9+ 默认）** | 把堆分成 2048 个 Region，可预测停顿，兼顾吞吐与延迟，内存大（>4G/6G）推荐 |
| **ZGC / Shenandoah**（JDK15/21） | 超低停顿（<1ms），用读屏障+染色指针，适合超大堆（TB 级） |

## 关键概念
- **STW（Stop-The-World）**：GC 时暂停所有业务线程。
- **GC Roots**：可作为根的对象（局部变量、静态变量、常量、JNI 引用等），可达性分析从根出发，不可达对象即被回收。
- **Full GC**：回收整个堆，停顿长，应尽量避免。
- **晋升（Promotion）**：年轻代对象存活超过阈值或 S 区放不下时，晋升到老年代。

## 常用调优参数
```
-Xms / -Xmx       堆初始/最大大小
-Xmn              年轻代大小
-XX:+UseG1GC      使用 G1
-XX:MaxGCPauseMillis=200   G1 期望停顿
```

**选择逻辑**：小堆追求吞吐 → Parallel；大堆追求低延迟 → G1 或 ZGC。

需要我深入讲某部分吗（比如 G1 的 Region 细节、GC 日志分析、调优实战）？

---

