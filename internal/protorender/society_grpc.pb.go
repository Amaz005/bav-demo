// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: proto/society.proto

package protorender

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion6

const (
	SocietyService_GetAllPostRPC_FullMethodName          = "/protorender.SocietyService/GetAllPostRPC"
	SocietyService_GetPostById_FullMethodName            = "/protorender.SocietyService/GetPostById"
	SocietyService_GetAllCommentsFromPost_FullMethodName = "/protorender.SocietyService/GetAllCommentsFromPost"
	SocietyService_UpdateComment_FullMethodName          = "/protorender.SocietyService/UpdateComment"
	SocietyService_AddComment_FullMethodName             = "/protorender.SocietyService/AddComment"
	SocietyService_AddPost_FullMethodName                = "/protorender.SocietyService/AddPost"
)

// SocietyServiceClient is the client API for SocietyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocietyServiceClient interface {
	GetAllPostRPC(ctx context.Context, in *RequestPosts, opts ...grpc.CallOption) (*ResponsePosts, error)
	GetPostById(ctx context.Context, in *RequestPost, opts ...grpc.CallOption) (*ResponsePost, error)
	GetAllCommentsFromPost(ctx context.Context, in *RequestPost, opts ...grpc.CallOption) (*ResponseComments, error)
	UpdateComment(ctx context.Context, in *RequestUpdateComment, opts ...grpc.CallOption) (*ResponseUpdate, error)
	AddComment(ctx context.Context, in *RequestAddComment, opts ...grpc.CallOption) (*ResponseUpdate, error)
	AddPost(ctx context.Context, in *RequestAddPost, opts ...grpc.CallOption) (*ResponseUpdate, error)
}

type societyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocietyServiceClient(cc grpc.ClientConnInterface) SocietyServiceClient {
	return &societyServiceClient{cc}
}

func (c *societyServiceClient) GetAllPostRPC(ctx context.Context, in *RequestPosts, opts ...grpc.CallOption) (*ResponsePosts, error) {
	out := new(ResponsePosts)
	err := c.cc.Invoke(ctx, SocietyService_GetAllPostRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) GetPostById(ctx context.Context, in *RequestPost, opts ...grpc.CallOption) (*ResponsePost, error) {
	out := new(ResponsePost)
	err := c.cc.Invoke(ctx, SocietyService_GetPostById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) GetAllCommentsFromPost(ctx context.Context, in *RequestPost, opts ...grpc.CallOption) (*ResponseComments, error) {
	out := new(ResponseComments)
	err := c.cc.Invoke(ctx, SocietyService_GetAllCommentsFromPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) UpdateComment(ctx context.Context, in *RequestUpdateComment, opts ...grpc.CallOption) (*ResponseUpdate, error) {
	out := new(ResponseUpdate)
	err := c.cc.Invoke(ctx, SocietyService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) AddComment(ctx context.Context, in *RequestAddComment, opts ...grpc.CallOption) (*ResponseUpdate, error) {
	out := new(ResponseUpdate)
	err := c.cc.Invoke(ctx, SocietyService_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) AddPost(ctx context.Context, in *RequestAddPost, opts ...grpc.CallOption) (*ResponseUpdate, error) {
	out := new(ResponseUpdate)
	err := c.cc.Invoke(ctx, SocietyService_AddPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocietyServiceServer is the server API for SocietyService service.
// All implementations must embed UnimplementedSocietyServiceServer
// for forward compatibility
type SocietyServiceServer interface {
	GetAllPostRPC(context.Context, *RequestPosts) (*ResponsePosts, error)
	GetPostById(context.Context, *RequestPost) (*ResponsePost, error)
	GetAllCommentsFromPost(context.Context, *RequestPost) (*ResponseComments, error)
	UpdateComment(context.Context, *RequestUpdateComment) (*ResponseUpdate, error)
	AddComment(context.Context, *RequestAddComment) (*ResponseUpdate, error)
	AddPost(context.Context, *RequestAddPost) (*ResponseUpdate, error)
	mustEmbedUnimplementedSocietyServiceServer()
}

// UnimplementedSocietyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSocietyServiceServer struct {
}

func (UnimplementedSocietyServiceServer) GetAllPostRPC(context.Context, *RequestPosts) (*ResponsePosts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPostRPC not implemented")
}
func (UnimplementedSocietyServiceServer) GetPostById(context.Context, *RequestPost) (*ResponsePost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedSocietyServiceServer) GetAllCommentsFromPost(context.Context, *RequestPost) (*ResponseComments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsFromPost not implemented")
}
func (UnimplementedSocietyServiceServer) UpdateComment(context.Context, *RequestUpdateComment) (*ResponseUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedSocietyServiceServer) AddComment(context.Context, *RequestAddComment) (*ResponseUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedSocietyServiceServer) AddPost(context.Context, *RequestAddPost) (*ResponseUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPost not implemented")
}
func (UnimplementedSocietyServiceServer) mustEmbedUnimplementedSocietyServiceServer() {}

// UnsafeSocietyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocietyServiceServer will
// result in compilation errors.
type UnsafeSocietyServiceServer interface {
	mustEmbedUnimplementedSocietyServiceServer()
}

func RegisterSocietyServiceServer(s grpc.Server, srv SocietyServiceServer) {
	s.RegisterService(&SocietyService_ServiceDesc, srv)
}

func _SocietyService_GetAllPostRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPosts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).GetAllPostRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_GetAllPostRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).GetAllPostRPC(ctx, req.(*RequestPosts))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_GetPostById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).GetPostById(ctx, req.(*RequestPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_GetAllCommentsFromPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).GetAllCommentsFromPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_GetAllCommentsFromPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).GetAllCommentsFromPost(ctx, req.(*RequestPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).UpdateComment(ctx, req.(*RequestUpdateComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).AddComment(ctx, req.(*RequestAddComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_AddPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).AddPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SocietyService_AddPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).AddPost(ctx, req.(*RequestAddPost))
	}
	return interceptor(ctx, in, info, handler)
}

// SocietyService_ServiceDesc is the grpc.ServiceDesc for SocietyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocietyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protorender.SocietyService",
	HandlerType: (*SocietyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPostRPC",
			Handler:    _SocietyService_GetAllPostRPC_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _SocietyService_GetPostById_Handler,
		},
		{
			MethodName: "GetAllCommentsFromPost",
			Handler:    _SocietyService_GetAllCommentsFromPost_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _SocietyService_UpdateComment_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _SocietyService_AddComment_Handler,
		},
		{
			MethodName: "AddPost",
			Handler:    _SocietyService_AddPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/society.proto",
}
