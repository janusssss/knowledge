### 概述

​	`iptables` 是Linux系统中用于配置IPv4数据包过滤和NAT的命令行工具。它是netfilter防火墙框架的用户空间接口。 

- 功能：网络数据包过滤系统
- 作用：防火墙、NAT、数据包修改
- 层级：用户空间命令 → 内核netfilter框架

### 结构 

- 四个表（Tables） 
  1. `filter`：数据包过滤（默认表）
  2. `nat`：网络地址转换
  3. `mangle`：数据包修改
  4. `raw`：数据包状态跟踪

- 三个链（Chains） 
  1. INPUT：进入本机的数据包
  2. OUTPUT：从本机发出的数据包
  3. FORWARD：经过本机转发的数据包

```bash
# 查看规则端口
iptables -L -n | grep 5432
# 新增链规则
iptables -I INPUT -p tcp -s 10.142.100.200 -d 10.142.149.168 --dport 5432 -j ACCEPT
# 删除链规则
iptables -I INPUT -p tcp -s 10.142.101.128 -d 10.142.149.168 --dport 5432 -j ACCEPT
```

