package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
	"userapi/configs"
	"userapi/container"
	grpc2 "userapi/grpc"
	"userapi/server"
)

func main() {
	var http, grpcServer, grpcClient bool
	flag.BoolVar(&http, "http", false, "run HTTP JSON API")
	flag.BoolVar(&grpcServer, "grpc-server", false, "run GRPC API")
	flag.BoolVar(&grpcClient, "grpc-client", false, "tests GRPC API")
	flag.Parse()

	logging := logrus.New()
	logging.SetReportCaller(true)

	config, err := configs.NewConfig()

	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := container.NewContainer(config, logging)

	if http {
		logrus.Info("http server starting")
		err = server.Run(config.Port, config.RedisPort, *containerInstance)
		if err != nil {
			logrus.Fatalf("error occured while running http server: %s, address: %s", err.Error(), config.Port)
		}
	} else if grpcServer {
		logrus.Info("grpc server starting")
		port := flag.Int("port", 50051, "The server port")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			logging.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		grpc2.RegisterUserServiceServer(s, grpc2.Server{containerInstance})
		logging.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logging.Fatalf("failed to serve: %v", err)
		}
	} else if grpcClient {
		addr := flag.String("addr", "localhost:50051", "the address to connect to")

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logging.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := grpc2.NewUserServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetAll(ctx, &grpc2.GetAllRequest{Page: 1})
		if err != nil {
			logging.Fatalf("could not connect: %v", err)
		}
		logging.Printf("Users list: %s", r.Users)
	}
}
