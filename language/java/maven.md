### mvn

```bash
# 命令行启动
mvn spring-boot:run -Dspring-boot.run.profiles=dev

# 运行主类
mvn exec:java
# 快速运行
mvn spring-boot:run -Dspring-boot.run.profiles=dev -Dmaven.main.skip=true -Dmaven.test.skip=true

# 编译源代码和测试代码
mvn compile test-compile
# 运行测试类
mvn test -Dtest=MyTest#myTestMethod # MyTest是你的测试类名，而myTestMethod是你想要运行的具体测试方法名
# 跳过编译编译阶段，直接运行测试函数
mvn surefire:test -Dtest=DacpServiceTest#Test

# 跳过测试代码的编译
mvn clean package -DskipTests

# 跳过依赖代表
mvn clean package -DskipTests -Dmaven.main.skip=true
mvn clean package -DskipTests -Dmaven.test.skip=true -Dmaven.main.skip=true

# 解决依赖
mvn dependency:resolve

# 提前下载项目所需的所有依赖（包括插件）
mvn dependency:go-offline
```



```bash
# 启动项目
mvn spring-boot:run
mvn spring-boot:run -Dspring-boot.run.profiles=janus

# 运行测试方法
mvn test -Dtest=SupplyManagementControllerTest#typeList
mvn test -Dtest=SupplyManagementControllerTest#typeList -DskipTests=false -Dmaven.test.skip=false
```

# 导入依赖
```bash
mvn install:install-file -Dfile=ojdbc6.jar -DgroupId=com.asiainfo.dacp -DartifactId=dacp-metamodel-flow -Dversion=1.5.0 -Dpackaging=jar -DgeneratePom=true
```

