package main

import (
	"log"
	"net"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	if err != nil {
		log.Fatal(err)
	}

	in, err := net.InterfaceByName("wlan0") if err != nil {
		log.Fatal(err)
	}

	sl := &syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_ALL,
		Ifindex:  in.Index,
	}

	if err := syscall.Bind(fd, sl); err != nil {
		log.Fatal(err)
	}

	frame := make([]byte, 0)
	destmac, err := net.ParseMAC("c8:58:c0:26:c5:41")
	frame = append(frame, destmac...)
	soumac, err := net.ParseMAC("88:f4:da:37:01:40")
	frame = append(frame, soumac...)
	frame = append(frame, []byte{0x13, 0x21}...)
	frame = append(frame, []byte("hello")...)

	if err := syscall.Sendto(fd, frame, 0, sl); err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}
