# 概述
go 是 Go 语言的官方命令行工具，用于管理 Go 源代码的编译、测试、安装等操作。

## 环境和工作空间管理
```bash
# 查看 Go 环境配置
go env

# 设置环境变量（临时）
go env -w GOPROXY=https://goproxy.cn,direct

# 查看 Go 版本
go version

# 下载依赖到缓存
go mod download
```

## 代码编译和运行
```bash
# 运行 Go 程序
go run main.go

# 运行多个文件
go run main.go utils.go

# 编译生成可执行文件
go build -o myapp main.go

# 编译当前包
go build .

# 编译并安装到 $GOPATH/bin
go install


# 禁用优化和内联（用于调试）
go build -gcflags="-N -l" main.go

# 显示编译详细信息
go build -x main.go

# 设置版本信息
go build -ldflags="-X main.version=1.0.0 -X main.buildTime=$(date)"
```

## 模块和依赖管理
```bash
# 初始化新模块
go mod init myproject

# 整理和优化 go.mod
go mod tidy

# 查看依赖图
go mod graph

# 添加缺失的依赖，移除未使用的依赖
go mod tidy

# 下载所有依赖
go mod download

# 设置国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 关闭代理校验（不推荐生产环境）
go env -w GOSUMDB=off
```

## 测试相关
```bash
# 运行测试
go test

# 运行测试并显示详细信息
go test -v

# 运行基准测试
go test -bench=.

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# 运行竞争检测
go test -race
```

## 代码质量和文档
```bash
# 静态代码分析
go vet .

# 代码格式化
go fmt ./...

# 生成文档并在浏览器查看
godoc -http=:6060

# 下载并安装第三方工具
go install golang.org/x/tools/cmd/goimports@latest
```

## 交叉编译
```bash
# 编译为 Linux 可执行文件
GOOS=linux GOARCH=amd64 go build -o app-linux main.go

# 编译为 Windows 可执行文件
GOOS=windows GOARCH=amd64 go build -o app.exe main.go

# 编译为 macOS 可执行文件
GOOS=darwin GOARCH=amd64 go build -o app-macos main.go
```