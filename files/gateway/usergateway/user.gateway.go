package usergateway

import (
	"common/discovery"
	pb "common/grpcs/user/pb"
	"context"
)

type UserGw struct {
	registry *discovery.Registry
}

func NewUserGw(registery *discovery.Registry) *UserGw {
	return &UserGw{
		registry: registery,
	}
}

func (u *UserGw) Guard(ctx context.Context, restructions []string) (*pb.GuardRes, error) {
	conn, err := u.registry.GetConn(ctx, "users")
	if err != nil {
		return nil, err
	}

	authSvcs := pb.NewAuthServiceClient(conn)

	res, err := authSvcs.Guard(ctx, &pb.GuardReq{
		Restrictions: restructions,
	})

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (u *UserGw) GetAllRoles(ctx context.Context) {

}
