package interceptor

import (
	"context"
	"log"
	"users-microservice/pkg/service"
	"users-microservice/util"

	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var MethCapas = map[string][]string{
	//users
	"/UserService/CreateUser": {"usersCreate"},
	"/UserService/ReadUser":   {"usersRead"},
	"/UserService/UpdateUser": {"usersUpdate"},
	"/UserService/DeleteUser": {"usersDelete"},
	"/UserService/Me":         {"auth"},
	"/UserService/ListUser":   {"auth"},
	//roles
	//capabilities
}

type ReqUser struct{}

type Intceptor struct {
	authsvcs *service.AuthSvcs
}

func NewInterceptor(i *do.Injector) *Intceptor {
	return &Intceptor{
		authsvcs: do.MustInvoke[*service.AuthSvcs](i),
	}
}

func (i *Intceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		caps := MethCapas[info.FullMethod]

		if len(caps) == 0 {
			//no capabilites required -proceed
			return handler(ctx, req)
		} else {
			authStatus, err := i.authsvcs.IsUserCan(ctx, caps...)
			if err != nil {
				return nil, status.Errorf(codes.PermissionDenied, "unauthorized RPC request rejected")
			}

			if authStatus.IsAuthorized {
				//TODO: add user to ctx
				newCtx := context.WithValue(ctx, util.Requser, authStatus.UserId)

				return handler(newCtx, req)
			} else {
				return nil, status.Errorf(codes.PermissionDenied, "unauthorized RPC request rejected")
			}
		}

	}
}
func (i *Intceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		caps := MethCapas[info.FullMethod]

		ctx := stream.Context()

		if len(caps) == 0 {
			//no capabilites required -proceed
			return handler(srv, stream)
		} else {
			authStatus, err := i.authsvcs.IsUserCan(ctx, caps...)
			if err != nil {
				return status.Errorf(codes.Unauthenticated, "not authenticated")
			}

			if authStatus.IsAuthorized {
				//TODO: add user to ctx
				//newCtx := context.WithValue(ctx, util.Requser, authStatus.UserId)

				return handler(srv, stream)
			} else {
				return status.Errorf(codes.Unauthenticated, "not authenticated")
			}
		}

	}
}
