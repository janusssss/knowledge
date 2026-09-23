# GRPC使用

## 步骤

1. go 环境。

2. 安装 protobuf 编译器。

   ```bash
   $ apt install -y protobuf-compiler
   $ protoc --version  # Ensure compiler version is 3+
   ```

3. 安装 go 插件。

   ```bash
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

4. 编译 proto 文件。

   ```bash
   protoc --go_out=. --go-grpc_out=. *.proto
   ```

   
