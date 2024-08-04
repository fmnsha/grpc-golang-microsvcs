// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: role_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRes, error)
	ReadRole(ctx context.Context, in *ReadRoleReq, opts ...grpc.CallOption) (*ReadRoleRes, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRes, error)
	DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRes, error)
	ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (RoleService_ListRoleClient, error)
	AddUserRole(ctx context.Context, in *AddUserRoleReq, opts ...grpc.CallOption) (*AddUserRoleRes, error)
	DeleteUserRole(ctx context.Context, in *DeleteUserRoleReq, opts ...grpc.CallOption) (*DeleteUserRoleRes, error)
	ListUserRoles(ctx context.Context, in *ListUserRolesReq, opts ...grpc.CallOption) (*ListUserRolesRes, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRes, error) {
	out := new(CreateRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) ReadRole(ctx context.Context, in *ReadRoleReq, opts ...grpc.CallOption) (*ReadRoleRes, error) {
	out := new(ReadRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/ReadRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRes, error) {
	out := new(UpdateRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRes, error) {
	out := new(DeleteRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (RoleService_ListRoleClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoleService_ServiceDesc.Streams[0], "/RoleService/ListRole", opts...)
	if err != nil {
		return nil, err
	}
	x := &roleServiceListRoleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoleService_ListRoleClient interface {
	Recv() (*ListRoleRes, error)
	grpc.ClientStream
}

type roleServiceListRoleClient struct {
	grpc.ClientStream
}

func (x *roleServiceListRoleClient) Recv() (*ListRoleRes, error) {
	m := new(ListRoleRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roleServiceClient) AddUserRole(ctx context.Context, in *AddUserRoleReq, opts ...grpc.CallOption) (*AddUserRoleRes, error) {
	out := new(AddUserRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/AddUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteUserRole(ctx context.Context, in *DeleteUserRoleReq, opts ...grpc.CallOption) (*DeleteUserRoleRes, error) {
	out := new(DeleteUserRoleRes)
	err := c.cc.Invoke(ctx, "/RoleService/DeleteUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) ListUserRoles(ctx context.Context, in *ListUserRolesReq, opts ...grpc.CallOption) (*ListUserRolesRes, error) {
	out := new(ListUserRolesRes)
	err := c.cc.Invoke(ctx, "/RoleService/ListUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRes, error)
	ReadRole(context.Context, *ReadRoleReq) (*ReadRoleRes, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRes, error)
	DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleRes, error)
	ListRole(*ListRoleReq, RoleService_ListRoleServer) error
	AddUserRole(context.Context, *AddUserRoleReq) (*AddUserRoleRes, error)
	DeleteUserRole(context.Context, *DeleteUserRoleReq) (*DeleteUserRoleRes, error)
	ListUserRoles(context.Context, *ListUserRolesReq) (*ListUserRolesRes, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) ReadRole(context.Context, *ReadRoleReq) (*ReadRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRole(context.Context, *DeleteRoleReq) (*DeleteRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleServiceServer) ListRole(*ListRoleReq, RoleService_ListRoleServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedRoleServiceServer) AddUserRole(context.Context, *AddUserRoleReq) (*AddUserRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteUserRole(context.Context, *DeleteUserRoleReq) (*DeleteUserRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRole not implemented")
}
func (UnimplementedRoleServiceServer) ListUserRoles(context.Context, *ListUserRolesReq) (*ListUserRolesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRoles not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_ReadRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).ReadRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/ReadRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).ReadRole(ctx, req.(*ReadRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRole(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_ListRole_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRoleReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoleServiceServer).ListRole(m, &roleServiceListRoleServer{stream})
}

type RoleService_ListRoleServer interface {
	Send(*ListRoleRes) error
	grpc.ServerStream
}

type roleServiceListRoleServer struct {
	grpc.ServerStream
}

func (x *roleServiceListRoleServer) Send(m *ListRoleRes) error {
	return x.ServerStream.SendMsg(m)
}

func _RoleService_AddUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).AddUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/AddUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).AddUserRole(ctx, req.(*AddUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/DeleteUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteUserRole(ctx, req.(*DeleteUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_ListUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRolesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).ListUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/ListUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).ListUserRoles(ctx, req.(*ListUserRolesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "ReadRole",
			Handler:    _RoleService_ReadRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleService_DeleteRole_Handler,
		},
		{
			MethodName: "AddUserRole",
			Handler:    _RoleService_AddUserRole_Handler,
		},
		{
			MethodName: "DeleteUserRole",
			Handler:    _RoleService_DeleteUserRole_Handler,
		},
		{
			MethodName: "ListUserRoles",
			Handler:    _RoleService_ListUserRoles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRole",
			Handler:       _RoleService_ListRole_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "role_service.proto",
}