package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		log.Fatal(err)
	}

	sa := syscall.SockaddrInet4{
		Port: 12358,
	}
	err = syscall.Bind(fd, &sa)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Listen(fd, 128)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("start end")

	connFd, _, err := syscall.Accept(fd)
	if err != nil {
		log.Fatal(err)
	}
	err = syscall.SetNonblock(connFd, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("accept fd:", connFd)
	buff := make([]byte, 1<<10)
	for {
		n, err := syscall.Read(connFd, buff)
		if err != nil {
			fmt.Println(err)
			if err == syscall.EAGAIN {
				time.Sleep(time.Second)
			}
			continue
		}
		if n <= 0 {
			fmt.Println("done")
			return
		}
		fmt.Println(string(buff[:n]))
	}
}
