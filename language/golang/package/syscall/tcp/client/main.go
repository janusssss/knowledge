package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":12358")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
	select {}
}
