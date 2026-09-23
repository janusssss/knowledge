# Web开发
- Gin
- Echo

# 数据库操作
- GORM
- go-redis

# 配置管理
- Viper: 完整的配置解决方案。支持从JSON、TOML、YAML、HCL、envfile、环境变量、命令行标志等读取配置。功能强大，是许多知名项目（如Hugo、Docker）的选择
- Cobra：一个用于创建强大的现代CLI应用程序的库。许多著名项目（如Docker、Kubernetes、Hugo）都用它来构建命令行工具。它通常与Viper结合使用。

# 测试与Mock
- Testify：最流行的测试工具包。它提供了更强大的断言（assert）、模拟（mock）和测试套件（suite）功能，极大地提升了Go测试的体验。
- gomock：Go官方提供的Mock框架，用于生成接口的模拟实现，配合go generate使用。

# 日志
- lumberjack: 是一个用于 日志滚动 的 Go 语言库。它是一个 io.Writer 的实现，这意味着它可以与任何需要 io.Writer 接口的日志库（如标准库 log、zap、logrus 等）无缝集成。
