// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package register

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterServiceClient interface {
	RegisterCustomer(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) RegisterCustomer(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/register.registerService/RegisterCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
// All implementations must embed UnimplementedRegisterServiceServer
// for forward compatibility
type RegisterServiceServer interface {
	RegisterCustomer(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedRegisterServiceServer()
}

// UnimplementedRegisterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (UnimplementedRegisterServiceServer) RegisterCustomer(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCustomer not implemented")
}
func (UnimplementedRegisterServiceServer) mustEmbedUnimplementedRegisterServiceServer() {}

// UnsafeRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServiceServer will
// result in compilation errors.
type UnsafeRegisterServiceServer interface {
	mustEmbedUnimplementedRegisterServiceServer()
}

func RegisterRegisterServiceServer(s grpc.ServiceRegistrar, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_RegisterCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).RegisterCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.registerService/RegisterCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).RegisterCustomer(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "register.registerService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCustomer",
			Handler:    _RegisterService_RegisterCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/register/register.proto",
}
