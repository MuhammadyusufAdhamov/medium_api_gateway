// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: like_service.proto

package post_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	Create(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Like, error)
	Get(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*Like, error)
	GetAll(ctx context.Context, in *GetAllLikesRequest, opts ...grpc.CallOption) (*GetAllLikesResponse, error)
	Update(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Like, error)
	Delete(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) Create(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Like, error) {
	out := new(Like)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Get(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*Like, error) {
	out := new(Like)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) GetAll(ctx context.Context, in *GetAllLikesRequest, opts ...grpc.CallOption) (*GetAllLikesResponse, error) {
	out := new(GetAllLikesResponse)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Update(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Like, error) {
	out := new(Like)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Delete(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility
type LikeServiceServer interface {
	Create(context.Context, *Like) (*Like, error)
	Get(context.Context, *GetLikeRequest) (*Like, error)
	GetAll(context.Context, *GetAllLikesRequest) (*GetAllLikesResponse, error)
	Update(context.Context, *Like) (*Like, error)
	Delete(context.Context, *GetLikeRequest) (*empty.Empty, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServiceServer struct {
}

func (UnimplementedLikeServiceServer) Create(context.Context, *Like) (*Like, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLikeServiceServer) Get(context.Context, *GetLikeRequest) (*Like, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLikeServiceServer) GetAll(context.Context, *GetAllLikesRequest) (*GetAllLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedLikeServiceServer) Update(context.Context, *Like) (*Like, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLikeServiceServer) Delete(context.Context, *GetLikeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Create(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Get(ctx, req.(*GetLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).GetAll(ctx, req.(*GetAllLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Update(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Delete(ctx, req.(*GetLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LikeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LikeService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _LikeService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LikeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LikeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like_service.proto",
}