### RabbitMQ queues

| Queue              | 说明                                 |
| ------------------ | ------------------------------------ |
| "Hello World!"     | 普通消息，pc一对一                   |
| Work Queues        | 工作队列，pc一对多，轮询消费         |
| Publish/Subscribe  | pc一对多，相当于广播                 |
| Routing            | 根据routing key推送信息给不同的queue |
| Topics             | 根据topic模式匹配推送消息到不同queue，支持通配符 |
| RPC                | 远程过程调用，实现同步的请求-响应模式，客户端发送请求等待服务端返回结果 |
| Publisher Confirms |                                      |

### RabbitMQ streams

| stream          |      |
| --------------- | ---- |
| "Hello World!"  |      |
| Offset Tracking |      |

### protocol

| protocol   |                                                             |
| ---------- | ----------------------------------------------------------- |
| AMQP 0-9-1 | `Advanced Message Queuing Protocol`, RabbitMQ主要实现的版本 |
| AMQP 1.0   | 更现代的版本，被更多企业级消息中间件支持                    |

### 启动

```bash
#!/bin/sh
# 24.04.2 LTS (Noble Numbat)
sudo apt-get install curl gnupg apt-transport-https -y

## Team RabbitMQ's signing key
curl -1sLf "https://keys.openpgp.org/vks/v1/by-fingerprint/0A9AF2115F4687BD29803A206B73A36E6026DFCA" | sudo gpg --dearmor | sudo tee /usr/share/keyrings/com.rabbitmq.team.gpg > /dev/null

## Add apt repositories maintained by Team RabbitMQ
sudo tee /etc/apt/sources.list.d/rabbitmq.list <<EOF
## Modern Erlang/OTP releases
##
deb [arch=amd64 signed-by=/usr/share/keyrings/com.rabbitmq.team.gpg] https://deb1.rabbitmq.com/rabbitmq-erlang/ubuntu/noble noble main
deb [arch=amd64 signed-by=/usr/share/keyrings/com.rabbitmq.team.gpg] https://deb2.rabbitmq.com/rabbitmq-erlang/ubuntu/noble noble main

## Latest RabbitMQ releases
##
deb [arch=amd64 signed-by=/usr/share/keyrings/com.rabbitmq.team.gpg] https://deb1.rabbitmq.com/rabbitmq-server/ubuntu/noble noble main
deb [arch=amd64 signed-by=/usr/share/keyrings/com.rabbitmq.team.gpg] https://deb2.rabbitmq.com/rabbitmq-server/ubuntu/noble noble main
EOF

## Update package indices
sudo apt-get update -y

## Install Erlang packages
sudo apt-get install -y erlang-base \
                        erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets \
                        erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key \
                        erlang-runtime-tools erlang-snmp erlang-ssl \
                        erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl

## Install rabbitmq-server and its dependencies
sudo apt-get install rabbitmq-server -y --fix-missing
```



```bash
docker run -d --name rabbitmq-management \
  -p 4369:4369 \
  -p 5671:5671 \
  -p 5672:5672 \
  -p 15671:15671 \
  -p 15672:15672 \
  -p 15691:15691 \
  -p 15692:15692 \
  -p 25672:25672 \
  rabbitmq:4.1.3-management

# management带web管理页面，没有的不带
```

| 端口  | 说明                                                         |
| ----- | ------------------------------------------------------------ |
| 4369  | Erlang 端口映射器守护进程 (Port Mapper Daemon)。RabbitMQ 基于 Erlang，此端口用于 Erlang 节点发现。 |
| 5671  | AMQP 协议（0-9-1 或 1.0）通过 TLS/SSL 加密传输。提供安全的消息队列服务。 |
| 5672  | AMQP 协议（0-9-1 或 1.0）的默认未加密端口。客户端主要通过此端口连接 RabbitMQ 进行消息收发。 |
| 15671 | RabbitMQ Management Plugin 的 HTTPS (加密) 端口。用于安全地访问 Web 管理界面或 HTTP API。 |
| 15672 | RabbitMQ Management Plugin 的 HTTP (非加密) 端口。用于访问 Web 管理界面 或 HTTP API。默认：guest/guest |
| 15691 | Prometheus 指标端口。RabbitMQ 通过此端口暴露运行时指标，供 Prometheus 抓取以进行监控。 |
| 15692 | RabbitMQ HTTP API 端口。通常用于程序化访问管理功能和获取节点状态等信息。 |
| 25672 | Erlang 分布式通信端口。用于 RabbitMQ 集群节点间的内部通信和状态同步。 |

### api

```bash
curl -u guest:guest http://localhost:15672/api/queues | jq
```

### 命令

```bash
rabbitmqctl list_queues
docker exec -it rabbitmq-management rabbitmqctl list_queues

# 查看应答情况
sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
# 查看exchange list
sudo rabbitmqctl list_exchanges
# 查看绑定关系
sudo rabbitmqctl list_bindings
```

### exchange类型

 exchange types available: `direct`, `topic`, `headers` and `fanout`

| type    | 概念                                                   |
| ------- | ------------------------------------------------------ |
| direct  | 直连交换机，根据完全匹配的路由键(route key)来路由消息  |
| topic   | 主题交换机，根据路由键的模式匹配来路由消息，支持通配符 |
| headers | 头部交换机，根据消息头部属性而不是路由键来路由消息     |
| fanout  | 扇出交换机，将消息广播到所有绑定的队列，忽略路由键     |
