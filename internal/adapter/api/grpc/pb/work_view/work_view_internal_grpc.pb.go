// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: work_view/work_view_internal.proto

package v1

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

// WorkViewInternalServiceClient is the client API for WorkViewInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkViewInternalServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	ProjectGet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Project, error)
}

type workViewInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkViewInternalServiceClient(cc grpc.ClientConnInterface) WorkViewInternalServiceClient {
	return &workViewInternalServiceClient{cc}
}

func (c *workViewInternalServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/work_view.WorkViewInternalService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workViewInternalServiceClient) ProjectGet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/work_view.WorkViewInternalService/ProjectGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkViewInternalServiceServer is the server API for WorkViewInternalService service.
// All implementations must embed UnimplementedWorkViewInternalServiceServer
// for forward compatibility
type WorkViewInternalServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	ProjectGet(context.Context, *ID) (*Project, error)
	mustEmbedUnimplementedWorkViewInternalServiceServer()
}

// UnimplementedWorkViewInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkViewInternalServiceServer struct {
}

func (UnimplementedWorkViewInternalServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedWorkViewInternalServiceServer) ProjectGet(context.Context, *ID) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectGet not implemented")
}
func (UnimplementedWorkViewInternalServiceServer) mustEmbedUnimplementedWorkViewInternalServiceServer() {
}

// UnsafeWorkViewInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkViewInternalServiceServer will
// result in compilation errors.
type UnsafeWorkViewInternalServiceServer interface {
	mustEmbedUnimplementedWorkViewInternalServiceServer()
}

func RegisterWorkViewInternalServiceServer(s grpc.ServiceRegistrar, srv WorkViewInternalServiceServer) {
	s.RegisterService(&WorkViewInternalService_ServiceDesc, srv)
}

func _WorkViewInternalService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkViewInternalServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/work_view.WorkViewInternalService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkViewInternalServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkViewInternalService_ProjectGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkViewInternalServiceServer).ProjectGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/work_view.WorkViewInternalService/ProjectGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkViewInternalServiceServer).ProjectGet(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkViewInternalService_ServiceDesc is the grpc.ServiceDesc for WorkViewInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkViewInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "work_view.WorkViewInternalService",
	HandlerType: (*WorkViewInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _WorkViewInternalService_SayHello_Handler,
		},
		{
			MethodName: "ProjectGet",
			Handler:    _WorkViewInternalService_ProjectGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "work_view/work_view_internal.proto",
}