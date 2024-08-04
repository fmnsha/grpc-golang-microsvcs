package interceptor

import (
	"context"
	"files-microservice/gateway/usergateway"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var MethCapas = map[string][]string{
	"/FileService/Upload":  {"auth"},
	"/FileService/GetById": {"auth"},
}

type Interceptor struct {
	usergw *usergateway.UserGw
}

func NewInterceptor(usergw *usergateway.UserGw) *Interceptor {
	return &Interceptor{
		usergw: usergw,
	}
}

func (i Interceptor) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	caps := MethCapas[info.FullMethod]

	if len(caps) == 0 {
		return handler(ctx, req)
	} else {
		st, err := i.usergw.Guard(ctx, caps)

		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if st.Authstatus.IsAuthorized {
			return handler(ctx, req)
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}

	}

}
func (i Interceptor) Stream() {}
