package main

import (
	"context"
	"flag"
	"go-interview/arithmetic/lemon/pd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

const (
	defaultName = "Janus"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
		}
	}(conn)
	c := pd.NewLoginServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", "token")
	defer cancel()
	r, err := c.Login(ctx, &pd.LoginRequest{Username: *name})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("login result: %v", r.Result)
}
