package service

import (
	"common/grpcs/user/pb"
	"context"
	"fmt"
	"users-microservice/gateway"
	"users-microservice/pkg/repo"
	"users-microservice/util"

	"github.com/samber/do"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserSvcs struct {
	pb.UnimplementedUserServiceServer
	userrepo repo.UserRepo
	emailgw  *gateway.EmailGw
}

func NewUserSvcs(i *do.Injector) (*UserSvcs, error) {
	return &UserSvcs{
		userrepo: do.MustInvoke[repo.UserRepo](i),
		emailgw:  do.MustInvoke[*gateway.EmailGw](i),
	}, nil
}

func (u *UserSvcs) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	user := req.GetUser()

	result, err := u.userrepo.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//send email
	_, errSend := u.emailgw.SendEmail(ctx, "ddd")

	if errSend != nil {
		fmt.Println(errSend)
	}

	return &pb.CreateUserRes{
		User: result,
	}, nil
}

func (u *UserSvcs) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	id := req.GetId()

	// md, ok := metadata.FromIncomingContext(ctx)
	// tenant_id := md["tenant_id"][0]
	// fmt.Println("the tenant", md["tenant_id"][0], ok)

	result, err := u.userrepo.ReadUser(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ReadUserRes{
		User: result,
	}, nil

}
func (u *UserSvcs) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {

	//todo: fetch req user form metadata
	user := req.GetUser()

	err := u.userrepo.UpdateUser(ctx, user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UpdateUserRes{
		User: user,
	}, nil
}

func (u *UserSvcs) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	id := req.GetId()
	err := u.userrepo.DeleteUser(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteUserRes{
		Success: true,
	}, nil
}

func (u *UserSvcs) ListUser(req *pb.ListUserReq, stream pb.UserService_ListUserServer) error {

	err := u.userrepo.ListUser(stream.Context(), func(user *pb.User) error {
		if err := stream.Send(&pb.ListUserRes{
			User: user,
		}); err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (u *UserSvcs) Me(ctx context.Context, req *emptypb.Empty) (*pb.MeRes, error) {
	// get user form ctx
	userId, ok := ctx.Value(util.Requser).(uint64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "error get user")
	}

	user, err := u.userrepo.ReadUser(ctx, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get user")
	}

	return &pb.MeRes{
		User: user,
	}, nil
}
