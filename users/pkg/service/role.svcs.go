package service

import (
	"common/grpcs/user/pb"
	"context"
	"users-microservice/pkg/repo"

	"github.com/samber/do"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc
type RoleSvcs struct {
	pb.UnimplementedRoleServiceServer
	rolerepo repo.RoleRepo
}

func NewRoleSvcs(i *do.Injector) (*RoleSvcs, error) {
	return &RoleSvcs{
		rolerepo: do.MustInvoke[repo.RoleRepo](i),
	}, nil
}

func (r *RoleSvcs) CreateRole(ctx context.Context, req *pb.CreateRoleReq) (*pb.CreateRoleRes, error) {

	role := req.GetRole()

	result, err := r.rolerepo.CreateRole(ctx, role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateRoleRes{
		Role: result,
	}, nil

}

func (r *RoleSvcs) ReadRole(ctx context.Context, req *pb.ReadRoleReq) (*pb.ReadRoleRes, error) {
	id := req.GetId()
	role, err := r.rolerepo.ReadRole(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.ReadRoleRes{
		Role: role,
	}, nil
}

func (r *RoleSvcs) UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (*pb.UpdateRoleRes, error) {
	role := req.GetRole()
	err := r.rolerepo.UpdateRole(ctx, role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UpdateRoleRes{
		Role: role,
	}, nil
}

func (r *RoleSvcs) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (*pb.DeleteRoleRes, error) {
	id := req.GetId()
	err := r.rolerepo.DeleteRole(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteRoleRes{
		Success: true,
	}, nil
}

func (r *RoleSvcs) ListRole(req *pb.ListRoleReq, stream pb.RoleService_ListRoleServer) error {
	err := r.rolerepo.ListRole(stream.Context(), func(role *pb.Role) error {
		if err := stream.Send(&pb.ListRoleRes{
			Role: role,
		}); err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

// ///
func (r *RoleSvcs) AddUserRole(ctx context.Context, req *pb.AddUserRoleReq) (*pb.AddUserRoleRes, error) {
	userrole := req.GetUserRole()

	if err := r.rolerepo.AddUserRole(ctx, userrole); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.AddUserRoleRes{
		Success: true,
	}, nil
}

func (r *RoleSvcs) DeleteUserRole(ctx context.Context, req *pb.DeleteUserRoleReq) (*pb.DeleteUserRoleRes, error) {
	id := req.GetId()
	if err := r.rolerepo.DeleteUserRole(ctx, id); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteUserRoleRes{
		Success: true,
	}, nil
}

func (r *RoleSvcs) ListUserRoles(ctx context.Context, req *pb.ListUserRolesReq) (*pb.ListUserRolesRes, error) {
	return nil, nil
}
