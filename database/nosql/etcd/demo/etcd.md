	ETCD 是一个高可用的分布式键值存储系统，常用于分布式系统的配置管理和服务发现。它是 CoreOS 团队开发的一个开源项目，广泛应用于  Kubernetes、Docker Swarm 等容器编排系统中。ETCD 的设计目标是提供一个简单、可靠且高性能的分布式存储解决方案。

- etcdctl : 命令行工具
  - get <key> : 获取键值
  - put <key> <value> : 设置键值
  - del --from-key / : 删除所有键数据
  - 
  - get  "" --prefix  --keys-only / : 查看所有的key值

```bash
etcdctl get  "" --prefix  --keys-only
```



```bash
#az1
#AppInfo:
#  Username: "root"
#  Password: "2021CTyun!"
#  EndPoint: "10.8.73.44:2379"
#  BakEndPoint: "10.8.73.45:2379"
#  MaxCallSendMsgSize: 10485760000
#  MaxCallRecvMsgSize: 107374182400
#  KeepAliveTime: 30
#  KeepAliveTimeout: 10
#MongoConf:
#  EndPoint: "gostack-mongos-0.gostack-mongos.az1.svc.cluster.net:27017"
#  EndPoints:
#    - "gostack-mongos-0.gostack-mongos.az1.svc.cluster.net:27017"
#  Username: "root"
#  Password: "test"
#  Timeout: 300
#MongoConfOld:
#  EndPoint: "gostack-mongos-0.gostack-mongos.az1.svc.cluster.net:27017"
#  EndPoints:
#    - "gostack-mongos-0.gostack-mongos.az1.svc.cluster.net:27017"
#  Username: "root"
#  Password: "test"
#  Timeout: 300000
# vm-az1
#AppInfo:
#  Username: "root"
#  Password: "ZmjsS3SuPP7^P&zk"
#  EndPoint: "10.8.75.22:2379"
#  BakEndPoint: "10.8.75.22:42379"
#  MaxCallSendMsgSize: 10485760000
#  MaxCallRecvMsgSize: 107374182400
#  KeepAliveTime: 30
#  KeepAliveTimeout: 10
#MongoConf:
#  EndPoint: "gostack-mongos.vm-az1.svc.cluster.net:27017"
#  EndPoints:
#    - gostack-mongos-0.gostack-mongos.vm-az1.svc.cluster.net:27017
#    - gostack-mongos-1.gostack-mongos.vm-az1.svc.cluster.net:27017
#    - gostack-mongos-2.gostack-mongos.vm-az1.svc.cluster.net:27017
#  Username: "root"
#  Password: "test"
#  Timeout: 300
#  Worker: 10
#  BatchSize: 100

#  az2
#AppInfo:
# Username: "root"
# Password: "2021CTyun!"
# EndPoint: "10.8.92.59:2379"
# BakEndPoint: "10.8.92.59:2379"
# MaxCallSendMsgSize: 10485760000
# MaxCallRecvMsgSize: 107374182400
# KeepAliveTime: 30
# KeepAliveTimeout: 10
#MongoConf:
#  EndPoint: "gostack-mongos.az2.svc.cluster.net:27017"
#  EndPoints:
#    - gostack-mongos-0.gostack-mongos.az2.svc.cluster.net:27017
#    - gostack-mongos-1.gostack-mongos.az2.svc.cluster.net:27017
#    - gostack-mongos-2.gostack-mongos.az2.svc.cluster.net:27017
#  Username: "root"
#  Password: "test"
#  Timeout: 300000
#  Worker: 10
#  BatchSize: 100
#az1
#AppInfo:
#  Username: "root"
#  Password: "2021CTyun!"
#  EndPoint: "10.8.73.45:2379"
#  BakEndPoint: "10.8.73.45:2379"
#  MaxCallSendMsgSize: 10485760000
#  MaxCallRecvMsgSize: 107374182400
#  KeepAliveTime: 30
#  KeepAliveTimeout: 10
#MongoConf:
#  EndPoint: "gostack-mongos.az1.svc.cluster.net:27017"
#  EndPoints:
#    - "gostack-mongos-0.gostack-mongos.az1.svc.cluster.net:27017"
      #- "gostack-mongos-1.gostack-mongos.az1.svc.cluster.net:27017"
      #- "gostack-mongos-2.gostack-mongos.az1.svc.cluster.net:27017"
      #  Username: "root"
      #  Password: "test"
      #  Timeout: 300
      #  Worker: 10
      #  BatchSize: 100
