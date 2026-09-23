package main

import (
	"context"
	"fmt"
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

	ctx, cancel := context.WithCancel(context.Background())
	rChan := cli.Watch(ctx, "key")
	for w := range rChan {
		for _, ev := range w.Events {
			fmt.Printf("Type: %s, Key: %q, Value: %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	cancel()
}
