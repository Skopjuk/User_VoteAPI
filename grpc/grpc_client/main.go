package main

import (
	"context"
	"flag"
	"log"
	"time"
	grpc2 "userapi/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpc2.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAll(ctx, &grpc2.GetAllRequest{Page: 1})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	log.Printf("Users list: %s", r.Users)
}
