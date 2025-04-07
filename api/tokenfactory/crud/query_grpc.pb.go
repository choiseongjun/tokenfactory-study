// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tokenfactory/crud/query.proto

package crud

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
	Query_Params_FullMethodName      = "/tokenfactory.crud.Query/Params"
	Query_Post_FullMethodName        = "/tokenfactory.crud.Query/Post"
	Query_PostAll_FullMethodName     = "/tokenfactory.crud.Query/PostAll"
	Query_PostHistory_FullMethodName = "/tokenfactory.crud.Query/PostHistory"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Post items.
	Post(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error)
	PostAll(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error)
	// Queries all historical states of a Post
	PostHistory(ctx context.Context, in *QueryPostHistoryRequest, opts ...grpc.CallOption) (*QueryPostHistoryResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Post(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error) {
	out := new(QueryGetPostResponse)
	err := c.cc.Invoke(ctx, Query_Post_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PostAll(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error) {
	out := new(QueryAllPostResponse)
	err := c.cc.Invoke(ctx, Query_PostAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PostHistory(ctx context.Context, in *QueryPostHistoryRequest, opts ...grpc.CallOption) (*QueryPostHistoryResponse, error) {
	out := new(QueryPostHistoryResponse)
	err := c.cc.Invoke(ctx, Query_PostHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Post items.
	Post(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error)
	PostAll(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error)
	// Queries all historical states of a Post
	PostHistory(context.Context, *QueryPostHistoryRequest) (*QueryPostHistoryResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Post(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedQueryServer) PostAll(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAll not implemented")
}
func (UnimplementedQueryServer) PostHistory(context.Context, *QueryPostHistoryRequest) (*QueryPostHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostHistory not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Post(ctx, req.(*QueryGetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PostAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PostAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PostAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PostAll(ctx, req.(*QueryAllPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PostHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPostHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PostHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PostHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PostHistory(ctx, req.(*QueryPostHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenfactory.crud.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Query_Post_Handler,
		},
		{
			MethodName: "PostAll",
			Handler:    _Query_PostAll_Handler,
		},
		{
			MethodName: "PostHistory",
			Handler:    _Query_PostHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenfactory/crud/query.proto",
}
