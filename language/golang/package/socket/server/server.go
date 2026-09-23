package main

import (
	"log"
	"net"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htos(syscall.ETH_P_ALL)))
	if err != nil {
		log.Fatal(err)
	}

	in, err := net.InterfaceByName("enp0s13f0u1u4")
	if err != nil {
		log.Fatal(err)
	}

	sl := &syscall.SockaddrLinklayer{
		Protocol: htos(syscall.ETH_P_ALL),
		Ifindex:  in.Index,
	}

	if err := syscall.Bind(fd, sl); err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 1<<10)
	for {
		//log.Println("About to call Recvfrom...") // 添加这行
		n, _, err := syscall.Recvfrom(fd, buff, 0)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 || n > 78 {
			continue
		}

		log.Println(buff[:n])
	}
}

func htos(port uint16) uint16 {
	return (port << 8) | (port >> 8)
}
