// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: blog.proto

package proto

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogId, error)
	GetOneBlog(ctx context.Context, in *BlogId, opts ...grpc.CallOption) (*Blog, error)
	UpdateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteBlog(ctx context.Context, in *BlogId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllBlog(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (BlogService_GetAllBlogClient, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*BlogId, error) {
	out := new(BlogId)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetOneBlog(ctx context.Context, in *BlogId, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, "/blog.BlogService/GetOneBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *BlogId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetAllBlog(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (BlogService_GetAllBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlogService_ServiceDesc.Streams[0], "/blog.BlogService/GetAllBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceGetAllBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_GetAllBlogClient interface {
	Recv() (*Blog, error)
	grpc.ClientStream
}

type blogServiceGetAllBlogClient struct {
	grpc.ClientStream
}

func (x *blogServiceGetAllBlogClient) Recv() (*Blog, error) {
	m := new(Blog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	CreateBlog(context.Context, *Blog) (*BlogId, error)
	GetOneBlog(context.Context, *BlogId) (*Blog, error)
	UpdateBlog(context.Context, *Blog) (*emptypb.Empty, error)
	DeleteBlog(context.Context, *BlogId) (*emptypb.Empty, error)
	GetAllBlog(*emptypb.Empty, BlogService_GetAllBlogServer) error
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreateBlog(context.Context, *Blog) (*BlogId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServiceServer) GetOneBlog(context.Context, *BlogId) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBlog not implemented")
}
func (UnimplementedBlogServiceServer) UpdateBlog(context.Context, *Blog) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServiceServer) DeleteBlog(context.Context, *BlogId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServiceServer) GetAllBlog(*emptypb.Empty, BlogService_GetAllBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllBlog not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetOneBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetOneBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/GetOneBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetOneBlog(ctx, req.(*BlogId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*BlogId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetAllBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).GetAllBlog(m, &blogServiceGetAllBlogServer{stream})
}

type BlogService_GetAllBlogServer interface {
	Send(*Blog) error
	grpc.ServerStream
}

type blogServiceGetAllBlogServer struct {
	grpc.ServerStream
}

func (x *blogServiceGetAllBlogServer) Send(m *Blog) error {
	return x.ServerStream.SendMsg(m)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "GetOneBlog",
			Handler:    _BlogService_GetOneBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllBlog",
			Handler:       _BlogService_GetAllBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog.proto",
}