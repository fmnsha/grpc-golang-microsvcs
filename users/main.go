package main

import (
	"common/discovery"
	"common/grpcs/user/pb"
	"context"
	"fmt"
	"log"
	"net"
	"users-microservice/interceptor"
	"users-microservice/pkg/service"
	"users-microservice/provider"

	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do"
	"google.golang.org/grpc"
)

func main() {

	injector := do.New()

	addr := "0.0.0.0:3000"

	ctx := context.Background()

	registry, err := discovery.NewRegistry()
	if err != nil {
		log.Fatal(err)
	}
	if err := registry.Register(ctx, "users", addr); err != nil {
		log.Fatal(err)
	}
	defer registry.DeRegister(ctx, "users")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	//provide registry
	do.ProvideValue(injector, registry)

	//provide services and repos and gateways
	provider.ProvideRepo(injector)
	provider.ProvideService(injector)
	provider.ProvideGateway(injector)

	//interceptor
	interceptor := interceptor.NewInterceptor(injector)

	//services
	authsvcs := do.MustInvoke[*service.AuthSvcs](injector)
	usersvcs := do.MustInvoke[*service.UserSvcs](injector)
	rolesvcs := do.MustInvoke[*service.RoleSvcs](injector)
	capabilitysvcs := do.MustInvoke[*service.CapabilitySvcs](injector)

	unaryInts := grpc.ChainUnaryInterceptor(interceptor.Unary())
	streamInts := grpc.ChainStreamInterceptor(interceptor.Stream())

	s := grpc.NewServer(unaryInts, streamInts)

	pb.RegisterUserServiceServer(s, usersvcs)
	pb.RegisterRoleServiceServer(s, rolesvcs)
	pb.RegisterCapabilityServiceServer(s, capabilitysvcs)
	pb.RegisterAuthServiceServer(s, authsvcs)
	fmt.Println("start user service ")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
