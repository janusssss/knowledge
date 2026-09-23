# 概述
ss是Socket Statistics的缩写，它是一个功能强大、高效且用于转储套接字统计信息的现代工具。它可以用来替代经典的 netstat 命令。

# 操作
```bash
# 查看tcp连接，默认ESTAB状态
ss -t

# 指定监听状态
ss -t state listening

# 指定目标端口（客户端->服务端），源端是sport
ss -t state listening dprot
```