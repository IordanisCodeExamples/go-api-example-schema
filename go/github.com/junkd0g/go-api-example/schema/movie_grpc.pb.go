// movie.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: movie.proto

package schema

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
	MovieService_GetMovie_FullMethodName = "/movieapi.MovieService/GetMovie"
)

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	// RPC method to fetch a movie by its title.
	GetMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) GetMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, MovieService_GetMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	// RPC method to fetch a movie by its title.
	GetMovie(context.Context, *MovieRequest) (*MovieResponse, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) GetMovie(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieService_GetMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movieapi.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieService_GetMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}