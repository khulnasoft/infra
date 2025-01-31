// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: orchestrator.proto

package orchestrator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SandboxServiceClient is the client API for SandboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandboxServiceClient interface {
	Create(ctx context.Context, in *SandboxCreateRequest, opts ...grpc.CallOption) (*SandboxCreateResponse, error)
	Update(ctx context.Context, in *SandboxUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SandboxListResponse, error)
	Delete(ctx context.Context, in *SandboxDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Pause(ctx context.Context, in *SandboxPauseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCachedBuilds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SandboxListCachedBuildsResponse, error)
}

type sandboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSandboxServiceClient(cc grpc.ClientConnInterface) SandboxServiceClient {
	return &sandboxServiceClient{cc}
}

func (c *sandboxServiceClient) Create(ctx context.Context, in *SandboxCreateRequest, opts ...grpc.CallOption) (*SandboxCreateResponse, error) {
	out := new(SandboxCreateResponse)
	err := c.cc.Invoke(ctx, "/SandboxService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) Update(ctx context.Context, in *SandboxUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SandboxService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SandboxListResponse, error) {
	out := new(SandboxListResponse)
	err := c.cc.Invoke(ctx, "/SandboxService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) Delete(ctx context.Context, in *SandboxDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SandboxService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) Pause(ctx context.Context, in *SandboxPauseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SandboxService/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) ListCachedBuilds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SandboxListCachedBuildsResponse, error) {
	out := new(SandboxListCachedBuildsResponse)
	err := c.cc.Invoke(ctx, "/SandboxService/ListCachedBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandboxServiceServer is the server API for SandboxService service.
// All implementations must embed UnimplementedSandboxServiceServer
// for forward compatibility
type SandboxServiceServer interface {
	Create(context.Context, *SandboxCreateRequest) (*SandboxCreateResponse, error)
	Update(context.Context, *SandboxUpdateRequest) (*emptypb.Empty, error)
	List(context.Context, *emptypb.Empty) (*SandboxListResponse, error)
	Delete(context.Context, *SandboxDeleteRequest) (*emptypb.Empty, error)
	Pause(context.Context, *SandboxPauseRequest) (*emptypb.Empty, error)
	ListCachedBuilds(context.Context, *emptypb.Empty) (*SandboxListCachedBuildsResponse, error)
	mustEmbedUnimplementedSandboxServiceServer()
}

// UnimplementedSandboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSandboxServiceServer struct {
}

func (UnimplementedSandboxServiceServer) Create(context.Context, *SandboxCreateRequest) (*SandboxCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSandboxServiceServer) Update(context.Context, *SandboxUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSandboxServiceServer) List(context.Context, *emptypb.Empty) (*SandboxListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSandboxServiceServer) Delete(context.Context, *SandboxDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSandboxServiceServer) Pause(context.Context, *SandboxPauseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedSandboxServiceServer) ListCachedBuilds(context.Context, *emptypb.Empty) (*SandboxListCachedBuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCachedBuilds not implemented")
}
func (UnimplementedSandboxServiceServer) mustEmbedUnimplementedSandboxServiceServer() {}

// UnsafeSandboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandboxServiceServer will
// result in compilation errors.
type UnsafeSandboxServiceServer interface {
	mustEmbedUnimplementedSandboxServiceServer()
}

func RegisterSandboxServiceServer(s grpc.ServiceRegistrar, srv SandboxServiceServer) {
	s.RegisterService(&SandboxService_ServiceDesc, srv)
}

func _SandboxService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).Create(ctx, req.(*SandboxCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).Update(ctx, req.(*SandboxUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).Delete(ctx, req.(*SandboxDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxPauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).Pause(ctx, req.(*SandboxPauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_ListCachedBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).ListCachedBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SandboxService/ListCachedBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).ListCachedBuilds(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SandboxService_ServiceDesc is the grpc.ServiceDesc for SandboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SandboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SandboxService",
	HandlerType: (*SandboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SandboxService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SandboxService_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SandboxService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SandboxService_Delete_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _SandboxService_Pause_Handler,
		},
		{
			MethodName: "ListCachedBuilds",
			Handler:    _SandboxService_ListCachedBuilds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator.proto",
}
