package main

import (
	"common/grpcs/file/pb"
	"context"
	"files-microservice/gateway/usergateway"
	"files-microservice/interceptor"
	"files-microservice/pkg/service"
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"

	"google.golang.org/grpc"

	"common/discovery"
)

func main() {
	addr := "0.0.0.0:3001"

	ctx := context.Background()

	registry, err := discovery.NewRegistry()
	if err != nil {
		log.Fatal(err)
	}
	if err := registry.Register(ctx, "files", addr); err != nil {
		log.Fatal(err)
	}

	//todo: close etcd client
	defer registry.DeRegister(ctx, "files")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	//services
	fileSvcs := service.NewFileSvcs()

	//gateways
	userGw := usergateway.NewUserGw(registry)

	//interceptors
	interceptor := interceptor.NewInterceptor(userGw)
	unaryInts := grpc.ChainUnaryInterceptor(interceptor.Unary)

	server := grpc.NewServer(unaryInts)
	pb.RegisterFileServiceServer(server, fileSvcs)
	fmt.Println("start server")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
