# 跨平台安装
```bash
```
docker run --privileged --rm tonistiigi/binfmt --install arm64
# 或者
安装下面两个包
qemu-user-static
qemu-user-static-binfmt

# 安装archlinuxarm
docker run --platform linux/arm64 --name archlinuxarm -it \
        -v /home/janus/Projects:/root \
        bamboocz/archlinuxarm:base-devel
docker start -i archlinuxarm
```
```

### docker ps

```bash
docker ps --format "TEMPLATE"

docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Status}}"
```

| 字段名        | 说明                            |
| ------------- | ------------------------------- |
| `.ID`         | 容器完整 ID                     |
| `.Image`      | 镜像名称                        |
| `.Command`    | 启动命令                        |
| `.CreatedAt`  | 创建时间                        |
| `.RunningFor` | 运行时长                        |
| `.Ports`      | 暴露的端口                      |
| `.Status`     | 容器状态                        |
| `.Names`      | 容器名称                        |
| `.Labels`     | 所有标签（键值对)               |
| `.Label`      | 单个标签（如 {{.Label "env"}}） |
| `.Mounts`     | 挂载的卷                        |
| `.Networks`   | 连接的网络                      |

### 镜像列表

```bash
docker image ls --format "table {{.Repository}}\t{{.Tag}}"
```

