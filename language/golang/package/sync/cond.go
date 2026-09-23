package main

import (
	"fmt"
	"sync"
	"time"
)

type ItemQueue struct {
	items []int
	cond  *sync.Cond
}

func NewItemQueue() *ItemQueue {
	iq := &ItemQueue{
		items: make([]int, 0),
	}
	iq.cond = sync.NewCond(&sync.Mutex{})
	return iq
}

func (iq *ItemQueue) Add(item int) {
	iq.cond.L.Lock()
	defer iq.cond.L.Unlock()

	iq.items = append(iq.items, item)
	fmt.Println("Added item:", item)
	iq.cond.Signal() // 通知等待的goroutine有新项目添加了
}

func (iq *ItemQueue) Remove() int {
	iq.cond.L.Lock()
	defer iq.cond.L.Unlock()

	for len(iq.items) == 0 {
		iq.cond.Wait() // 如果没有项目可移除，则等待
	}

	item := iq.items[0]
	iq.items = iq.items[1:]
	fmt.Println("Removed item:", item)
	return item
}

func main() {
	iq := NewItemQueue()

	// 启动消费者goroutine
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second) // 模拟处理时间
			removedItem := iq.Remove()
			fmt.Println("Consumed:", removedItem)
		}
	}()

	// 生产者添加项目
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2) // 模拟生产间隔
		iq.Add(i)
	}
}
