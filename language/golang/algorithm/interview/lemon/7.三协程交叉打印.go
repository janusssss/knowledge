// 柠檬科技
package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，启动三个线程，三个线程的ID分别是A，B，C；，每个线程将自己的ID值在屏幕上打印5遍，打印顺序是ABCABC
func main() {
	jod, n := sync.WaitGroup{}, 3
	jod.Add(n)
	chs := make([]chan int, n)
	for i, _ := range [3]struct{}{} {
		chs[i] = make(chan int, 1)
		if i == 0 {
			chs[0] <- 0
		}
		go printName(i, &jod, chs)
	}
	jod.Wait()
}

func printName(i int, job *sync.WaitGroup, chs []chan int) {
	for range [5]struct{}{} {
		<-chs[i]
		fmt.Print(string(rune('A' + i)))
		t := (i + 1) % 3
		chs[t] <- 0
	}
	job.Done()
}
