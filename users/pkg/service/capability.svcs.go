package service

import (
	"common/grpcs/user/pb"
	"context"
	"users-microservice/pkg/repo"

	"github.com/samber/do"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CapabilitySvcs struct {
	pb.UnimplementedCapabilityServiceServer
	capabilityrepo repo.CapabilityRepo
}

func NewCapabilitySvcs(i *do.Injector) (*CapabilitySvcs, error) {
	return &CapabilitySvcs{
		capabilityrepo: do.MustInvoke[repo.CapabilityRepo](i),
	}, nil
}

func (c *CapabilitySvcs) CreateCap(ctx context.Context, req *pb.CreateCapReq) (*pb.CreateCapRes, error) {
	capability := req.GetCapability()
	err := c.capabilityrepo.CreateCap(ctx, capability)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreateCapRes{
		Capability: capability,
	}, nil

}
func (c *CapabilitySvcs) ReadCap(ctx context.Context, req *pb.ReadCapReq) (*pb.ReadCapRes, error) {
	id := req.GetId()
	capability, err := c.capabilityrepo.ReadCap(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ReadCapRes{
		Capability: capability,
	}, nil
}
func (c *CapabilitySvcs) UpdateCap(ctx context.Context, req *pb.UpdateCapReq) (*pb.UpdateCapRes, error) {
	capability := req.GetCapability()
	err := c.capabilityrepo.UpdateCap(ctx, capability)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UpdateCapRes{
		Capability: capability,
	}, nil
}
func (c *CapabilitySvcs) DeleteCap(ctx context.Context, req *pb.DeleteCapReq) (*pb.DeleteCapRes, error) {
	id := req.GetId()
	err := c.capabilityrepo.DeleteCap(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteCapRes{
		Success: true,
	}, nil
}

func (c *CapabilitySvcs) ListCap(req *pb.ListCapReq, stream pb.CapabilityService_ListCapServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCap not implemented")
}

func (c *CapabilitySvcs) AddRoleCapability(ctx context.Context, req *pb.AddRoleCapabilitiesReq) (*pb.AddRoleCapabilitiesRes, error) {
	roleCapability := req.GetRoleCapability()

	if err := c.capabilityrepo.AddRoleCapabilities(ctx, roleCapability); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.AddRoleCapabilitiesRes{
		Success: true,
	}, nil
}
