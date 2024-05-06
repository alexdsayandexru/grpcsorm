// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: sorm.proto

package sorm

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

// UserDataManagementClient is the client API for UserDataManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataManagementClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	LogoutUser(ctx context.Context, in *LogoutUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	DeleteUserAccount(ctx context.Context, in *DeleteUserAccountRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	UpdateUserData(ctx context.Context, in *UpdateUserDataRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	DeleteUserAccountAdmin(ctx context.Context, in *DeleteUserAccountAdminRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	UpdateUserDataAdmin(ctx context.Context, in *UpdateUserDataAdminRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	UserAccountRecovery(ctx context.Context, in *UserAccountRecoveryRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	DirectoryData(ctx context.Context, in *DirectoryDataRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
	MigrateUser(ctx context.Context, in *MigrateUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error)
}

type userDataManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataManagementClient(cc grpc.ClientConnInterface) UserDataManagementClient {
	return &userDataManagementClient{cc}
}

func (c *userDataManagementClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) LogoutUser(ctx context.Context, in *LogoutUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/LogoutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) DeleteUserAccount(ctx context.Context, in *DeleteUserAccountRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/DeleteUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) UpdateUserData(ctx context.Context, in *UpdateUserDataRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/UpdateUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) DeleteUserAccountAdmin(ctx context.Context, in *DeleteUserAccountAdminRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/DeleteUserAccountAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) UpdateUserDataAdmin(ctx context.Context, in *UpdateUserDataAdminRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/UpdateUserDataAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) UserAccountRecovery(ctx context.Context, in *UserAccountRecoveryRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/UserAccountRecovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) DirectoryData(ctx context.Context, in *DirectoryDataRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/DirectoryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataManagementClient) MigrateUser(ctx context.Context, in *MigrateUserRequest, opts ...grpc.CallOption) (*DataManagementResponse, error) {
	out := new(DataManagementResponse)
	err := c.cc.Invoke(ctx, "/sorm.UserDataManagement/MigrateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataManagementServer is the server API for UserDataManagement service.
// All implementations must embed UnimplementedUserDataManagementServer
// for forward compatibility
type UserDataManagementServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*DataManagementResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*DataManagementResponse, error)
	LogoutUser(context.Context, *LogoutUserRequest) (*DataManagementResponse, error)
	DeleteUserAccount(context.Context, *DeleteUserAccountRequest) (*DataManagementResponse, error)
	UpdateUserData(context.Context, *UpdateUserDataRequest) (*DataManagementResponse, error)
	DeleteUserAccountAdmin(context.Context, *DeleteUserAccountAdminRequest) (*DataManagementResponse, error)
	UpdateUserDataAdmin(context.Context, *UpdateUserDataAdminRequest) (*DataManagementResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DataManagementResponse, error)
	UserAccountRecovery(context.Context, *UserAccountRecoveryRequest) (*DataManagementResponse, error)
	DirectoryData(context.Context, *DirectoryDataRequest) (*DataManagementResponse, error)
	MigrateUser(context.Context, *MigrateUserRequest) (*DataManagementResponse, error)
	mustEmbedUnimplementedUserDataManagementServer()
}

// UnimplementedUserDataManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataManagementServer struct {
}

func (UnimplementedUserDataManagementServer) RegisterUser(context.Context, *RegisterUserRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserDataManagementServer) LoginUser(context.Context, *LoginUserRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserDataManagementServer) LogoutUser(context.Context, *LogoutUserRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedUserDataManagementServer) DeleteUserAccount(context.Context, *DeleteUserAccountRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccount not implemented")
}
func (UnimplementedUserDataManagementServer) UpdateUserData(context.Context, *UpdateUserDataRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserData not implemented")
}
func (UnimplementedUserDataManagementServer) DeleteUserAccountAdmin(context.Context, *DeleteUserAccountAdminRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccountAdmin not implemented")
}
func (UnimplementedUserDataManagementServer) UpdateUserDataAdmin(context.Context, *UpdateUserDataAdminRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDataAdmin not implemented")
}
func (UnimplementedUserDataManagementServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedUserDataManagementServer) UserAccountRecovery(context.Context, *UserAccountRecoveryRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAccountRecovery not implemented")
}
func (UnimplementedUserDataManagementServer) DirectoryData(context.Context, *DirectoryDataRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectoryData not implemented")
}
func (UnimplementedUserDataManagementServer) MigrateUser(context.Context, *MigrateUserRequest) (*DataManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateUser not implemented")
}
func (UnimplementedUserDataManagementServer) mustEmbedUnimplementedUserDataManagementServer() {}

// UnsafeUserDataManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataManagementServer will
// result in compilation errors.
type UnsafeUserDataManagementServer interface {
	mustEmbedUnimplementedUserDataManagementServer()
}

func RegisterUserDataManagementServer(s grpc.ServiceRegistrar, srv UserDataManagementServer) {
	s.RegisterService(&UserDataManagement_ServiceDesc, srv)
}

func _UserDataManagement_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/LogoutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).LogoutUser(ctx, req.(*LogoutUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_DeleteUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).DeleteUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/DeleteUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).DeleteUserAccount(ctx, req.(*DeleteUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_UpdateUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).UpdateUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/UpdateUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).UpdateUserData(ctx, req.(*UpdateUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_DeleteUserAccountAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAccountAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).DeleteUserAccountAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/DeleteUserAccountAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).DeleteUserAccountAdmin(ctx, req.(*DeleteUserAccountAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_UpdateUserDataAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDataAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).UpdateUserDataAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/UpdateUserDataAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).UpdateUserDataAdmin(ctx, req.(*UpdateUserDataAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_UserAccountRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccountRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).UserAccountRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/UserAccountRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).UserAccountRecovery(ctx, req.(*UserAccountRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_DirectoryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectoryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).DirectoryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/DirectoryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).DirectoryData(ctx, req.(*DirectoryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataManagement_MigrateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataManagementServer).MigrateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.UserDataManagement/MigrateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataManagementServer).MigrateUser(ctx, req.(*MigrateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataManagement_ServiceDesc is the grpc.ServiceDesc for UserDataManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sorm.UserDataManagement",
	HandlerType: (*UserDataManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserDataManagement_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserDataManagement_LoginUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _UserDataManagement_LogoutUser_Handler,
		},
		{
			MethodName: "DeleteUserAccount",
			Handler:    _UserDataManagement_DeleteUserAccount_Handler,
		},
		{
			MethodName: "UpdateUserData",
			Handler:    _UserDataManagement_UpdateUserData_Handler,
		},
		{
			MethodName: "DeleteUserAccountAdmin",
			Handler:    _UserDataManagement_DeleteUserAccountAdmin_Handler,
		},
		{
			MethodName: "UpdateUserDataAdmin",
			Handler:    _UserDataManagement_UpdateUserDataAdmin_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _UserDataManagement_DeleteAccount_Handler,
		},
		{
			MethodName: "UserAccountRecovery",
			Handler:    _UserDataManagement_UserAccountRecovery_Handler,
		},
		{
			MethodName: "DirectoryData",
			Handler:    _UserDataManagement_DirectoryData_Handler,
		},
		{
			MethodName: "MigrateUser",
			Handler:    _UserDataManagement_MigrateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorm.proto",
}