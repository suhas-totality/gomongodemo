// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/v1/proto/definition/service.proto

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

const (
	Service_CreateAccount_FullMethodName              = "/api.v1.gomongodemo.Service/CreateAccount"
	Service_ReadAccount_FullMethodName                = "/api.v1.gomongodemo.Service/ReadAccount"
	Service_ReadAccountByAccountNumber_FullMethodName = "/api.v1.gomongodemo.Service/ReadAccountByAccountNumber"
	Service_UpdateAccount_FullMethodName              = "/api.v1.gomongodemo.Service/UpdateAccount"
	Service_DeleteAcount_FullMethodName               = "/api.v1.gomongodemo.Service/DeleteAcount"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	ReadAccount(ctx context.Context, in *ReadAccountRequest, opts ...grpc.CallOption) (*ReadAccountResponse, error)
	ReadAccountByAccountNumber(ctx context.Context, in *ReadAccountByAccountNumberRequest, opts ...grpc.CallOption) (*ReadAccountByAccountNumberResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	DeleteAcount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, Service_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ReadAccount(ctx context.Context, in *ReadAccountRequest, opts ...grpc.CallOption) (*ReadAccountResponse, error) {
	out := new(ReadAccountResponse)
	err := c.cc.Invoke(ctx, Service_ReadAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ReadAccountByAccountNumber(ctx context.Context, in *ReadAccountByAccountNumberRequest, opts ...grpc.CallOption) (*ReadAccountByAccountNumberResponse, error) {
	out := new(ReadAccountByAccountNumberResponse)
	err := c.cc.Invoke(ctx, Service_ReadAccountByAccountNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, Service_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteAcount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, Service_DeleteAcount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	ReadAccount(context.Context, *ReadAccountRequest) (*ReadAccountResponse, error)
	ReadAccountByAccountNumber(context.Context, *ReadAccountByAccountNumberRequest) (*ReadAccountByAccountNumberResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	DeleteAcount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedServiceServer) ReadAccount(context.Context, *ReadAccountRequest) (*ReadAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAccount not implemented")
}
func (UnimplementedServiceServer) ReadAccountByAccountNumber(context.Context, *ReadAccountByAccountNumberRequest) (*ReadAccountByAccountNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAccountByAccountNumber not implemented")
}
func (UnimplementedServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedServiceServer) DeleteAcount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAcount not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ReadAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ReadAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ReadAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ReadAccount(ctx, req.(*ReadAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ReadAccountByAccountNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAccountByAccountNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ReadAccountByAccountNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ReadAccountByAccountNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ReadAccountByAccountNumber(ctx, req.(*ReadAccountByAccountNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteAcount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteAcount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteAcount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteAcount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.gomongodemo.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Service_CreateAccount_Handler,
		},
		{
			MethodName: "ReadAccount",
			Handler:    _Service_ReadAccount_Handler,
		},
		{
			MethodName: "ReadAccountByAccountNumber",
			Handler:    _Service_ReadAccountByAccountNumber_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Service_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAcount",
			Handler:    _Service_DeleteAcount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/proto/definition/service.proto",
}
