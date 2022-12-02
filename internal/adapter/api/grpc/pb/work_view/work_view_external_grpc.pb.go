// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: work_view/work_view_external.proto

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

// WorkViewExternalServiceClient is the client API for WorkViewExternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkViewExternalServiceClient interface {
	ProjectCreate(ctx context.Context, in *ProjectCreateRequest, opts ...grpc.CallOption) (*ID, error)
	ProjectGet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Project, error)
	ProjectDelete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type workViewExternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkViewExternalServiceClient(cc grpc.ClientConnInterface) WorkViewExternalServiceClient {
	return &workViewExternalServiceClient{cc}
}

func (c *workViewExternalServiceClient) ProjectCreate(ctx context.Context, in *ProjectCreateRequest, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/work_view.WorkViewExternalService/ProjectCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workViewExternalServiceClient) ProjectGet(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/work_view.WorkViewExternalService/ProjectGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workViewExternalServiceClient) ProjectDelete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/work_view.WorkViewExternalService/ProjectDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkViewExternalServiceServer is the server API for WorkViewExternalService service.
// All implementations must embed UnimplementedWorkViewExternalServiceServer
// for forward compatibility
type WorkViewExternalServiceServer interface {
	ProjectCreate(context.Context, *ProjectCreateRequest) (*ID, error)
	ProjectGet(context.Context, *ID) (*Project, error)
	ProjectDelete(context.Context, *ID) (*Empty, error)
	mustEmbedUnimplementedWorkViewExternalServiceServer()
}

// UnimplementedWorkViewExternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkViewExternalServiceServer struct {
}

func (UnimplementedWorkViewExternalServiceServer) ProjectCreate(context.Context, *ProjectCreateRequest) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectCreate not implemented")
}
func (UnimplementedWorkViewExternalServiceServer) ProjectGet(context.Context, *ID) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectGet not implemented")
}
func (UnimplementedWorkViewExternalServiceServer) ProjectDelete(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectDelete not implemented")
}
func (UnimplementedWorkViewExternalServiceServer) mustEmbedUnimplementedWorkViewExternalServiceServer() {
}

// UnsafeWorkViewExternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkViewExternalServiceServer will
// result in compilation errors.
type UnsafeWorkViewExternalServiceServer interface {
	mustEmbedUnimplementedWorkViewExternalServiceServer()
}

func RegisterWorkViewExternalServiceServer(s grpc.ServiceRegistrar, srv WorkViewExternalServiceServer) {
	s.RegisterService(&WorkViewExternalService_ServiceDesc, srv)
}

func _WorkViewExternalService_ProjectCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkViewExternalServiceServer).ProjectCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/work_view.WorkViewExternalService/ProjectCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkViewExternalServiceServer).ProjectCreate(ctx, req.(*ProjectCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkViewExternalService_ProjectGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkViewExternalServiceServer).ProjectGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/work_view.WorkViewExternalService/ProjectGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkViewExternalServiceServer).ProjectGet(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkViewExternalService_ProjectDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkViewExternalServiceServer).ProjectDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/work_view.WorkViewExternalService/ProjectDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkViewExternalServiceServer).ProjectDelete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkViewExternalService_ServiceDesc is the grpc.ServiceDesc for WorkViewExternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkViewExternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "work_view.WorkViewExternalService",
	HandlerType: (*WorkViewExternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProjectCreate",
			Handler:    _WorkViewExternalService_ProjectCreate_Handler,
		},
		{
			MethodName: "ProjectGet",
			Handler:    _WorkViewExternalService_ProjectGet_Handler,
		},
		{
			MethodName: "ProjectDelete",
			Handler:    _WorkViewExternalService_ProjectDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "work_view/work_view_external.proto",
}