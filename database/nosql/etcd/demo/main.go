package main

import (
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func main() {
	// 创建etcd客户端配置
	cfg := clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,            // 超时时间
	}

	// 建立连接
	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 使用客户端进行操作...
}
