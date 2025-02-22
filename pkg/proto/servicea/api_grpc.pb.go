// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: servicea/api.proto

package servicea

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

// ServiceAAPIClient is the client API for ServiceAAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceAAPIClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type serviceAAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAAPIClient(cc grpc.ClientConnInterface) ServiceAAPIClient {
	return &serviceAAPIClient{cc}
}

func (c *serviceAAPIClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/servicea.ServiceAAPI/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAAPIServer is the server API for ServiceAAPI service.
// All implementations must embed UnimplementedServiceAAPIServer
// for forward compatibility
type ServiceAAPIServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedServiceAAPIServer()
}

// UnimplementedServiceAAPIServer must be embedded to have forward compatible implementations.
type UnimplementedServiceAAPIServer struct {
}

func (UnimplementedServiceAAPIServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedServiceAAPIServer) mustEmbedUnimplementedServiceAAPIServer() {}

// UnsafeServiceAAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAAPIServer will
// result in compilation errors.
type UnsafeServiceAAPIServer interface {
	mustEmbedUnimplementedServiceAAPIServer()
}

func RegisterServiceAAPIServer(s grpc.ServiceRegistrar, srv ServiceAAPIServer) {
	s.RegisterService(&ServiceAAPI_ServiceDesc, srv)
}

func _ServiceAAPI_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAAPIServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicea.ServiceAAPI/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAAPIServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceAAPI_ServiceDesc is the grpc.ServiceDesc for ServiceAAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceAAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servicea.ServiceAAPI",
	HandlerType: (*ServiceAAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ServiceAAPI_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicea/api.proto",
}
