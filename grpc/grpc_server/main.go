package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"userapi/configs"
	"userapi/container"
	grpc2 "userapi/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	logging := logrus.New()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	config, err := configs.NewConfig()

	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := container.NewContainer(config, logging)
	s := grpc.NewServer()
	grpc2.RegisterUserServiceServer(s, grpc2.Server{containerInstance})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
