package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cfg := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cli.Delete(ctx, "key")
	if err != nil {
		panic(err)
	}
}
