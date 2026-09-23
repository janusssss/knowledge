package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	var maxConcurrency int64 = 8 // 设置最大并发数为2
	sem := semaphore.NewWeighted(maxConcurrency)

	ctx := context.Background()

	for i := 0; i < 5; i++ {
		go func(id int) {
			if err := sem.Acquire(ctx, 2); err != nil {
				log.Printf("Failed to acquire semaphore in goroutine %d: %v", id, err)
				return
			}
			defer sem.Release(1)

			fmt.Printf("Goroutine %d is executing...\n", id)
			time.Sleep(2 * time.Second) // 模拟耗时操作
			fmt.Printf("Goroutine %d has finished execution.\n", id)
		}(i)
	}

	// 等待所有goroutine完成
	time.Sleep(10 * time.Second)
}
