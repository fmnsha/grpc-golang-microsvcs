package main

import (
	"context"
	"emails-microservice/pkg/service"
	"fmt"
	"log"
	"net"

	"common/discovery"
	"common/grpcs/email/pb"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {

	addr := "0.0.0.0:3002"

	ctx := context.Background()

	registry, err := discovery.NewRegistry()
	if err != nil {
		log.Fatal(err)
	}
	if err := registry.Register(ctx, "emails", addr); err != nil {
		log.Fatal(err)
	}

	//todo: close etcd client
	defer registry.DeRegister(ctx, "emails")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	//services
	emailSvcs := service.NewEmailSvcs()

	//gateways

	//interceptors
	// interceptor := interceptor.NewInterceptor(userGw)
	// unaryInts := grpc.ChainUnaryInterceptor(interceptor.Unary)

	server := grpc.NewServer()
	pb.RegisterEmailServiceServer(server, emailSvcs)
	fmt.Println("start server")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
