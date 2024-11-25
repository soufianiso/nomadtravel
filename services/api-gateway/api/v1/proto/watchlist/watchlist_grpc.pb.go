// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/watchlist/watchlist.proto

package watchlist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Watchlist_ListWatchlist_FullMethodName = "/watchlist.Watchlist/ListWatchlist"
	Watchlist_AddMovie_FullMethodName      = "/watchlist.Watchlist/AddMovie"
	Watchlist_RemoveMovie_FullMethodName   = "/watchlist.Watchlist/RemoveMovie"
)

// WatchlistClient is the client API for Watchlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchlistClient interface {
	ListWatchlist(ctx context.Context, in *ListWatchlistRequest, opts ...grpc.CallOption) (*ListWatchlistResponse, error)
	AddMovie(ctx context.Context, in *AddMovieRequest, opts ...grpc.CallOption) (*AddMovieResponse, error)
	RemoveMovie(ctx context.Context, in *RemoveMovieRequest, opts ...grpc.CallOption) (*RemoveMovieResponse, error)
}

type watchlistClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchlistClient(cc grpc.ClientConnInterface) WatchlistClient {
	return &watchlistClient{cc}
}

func (c *watchlistClient) ListWatchlist(ctx context.Context, in *ListWatchlistRequest, opts ...grpc.CallOption) (*ListWatchlistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWatchlistResponse)
	err := c.cc.Invoke(ctx, Watchlist_ListWatchlist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchlistClient) AddMovie(ctx context.Context, in *AddMovieRequest, opts ...grpc.CallOption) (*AddMovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMovieResponse)
	err := c.cc.Invoke(ctx, Watchlist_AddMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watchlistClient) RemoveMovie(ctx context.Context, in *RemoveMovieRequest, opts ...grpc.CallOption) (*RemoveMovieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveMovieResponse)
	err := c.cc.Invoke(ctx, Watchlist_RemoveMovie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchlistServer is the server API for Watchlist service.
// All implementations must embed UnimplementedWatchlistServer
// for forward compatibility.
type WatchlistServer interface {
	ListWatchlist(context.Context, *ListWatchlistRequest) (*ListWatchlistResponse, error)
	AddMovie(context.Context, *AddMovieRequest) (*AddMovieResponse, error)
	RemoveMovie(context.Context, *RemoveMovieRequest) (*RemoveMovieResponse, error)
	mustEmbedUnimplementedWatchlistServer()
}

// UnimplementedWatchlistServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWatchlistServer struct{}

func (UnimplementedWatchlistServer) ListWatchlist(context.Context, *ListWatchlistRequest) (*ListWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWatchlist not implemented")
}
func (UnimplementedWatchlistServer) AddMovie(context.Context, *AddMovieRequest) (*AddMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedWatchlistServer) RemoveMovie(context.Context, *RemoveMovieRequest) (*RemoveMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMovie not implemented")
}
func (UnimplementedWatchlistServer) mustEmbedUnimplementedWatchlistServer() {}
func (UnimplementedWatchlistServer) testEmbeddedByValue()                   {}

// UnsafeWatchlistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchlistServer will
// result in compilation errors.
type UnsafeWatchlistServer interface {
	mustEmbedUnimplementedWatchlistServer()
}

func RegisterWatchlistServer(s grpc.ServiceRegistrar, srv WatchlistServer) {
	// If the following call pancis, it indicates UnimplementedWatchlistServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Watchlist_ServiceDesc, srv)
}

func _Watchlist_ListWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchlistServer).ListWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watchlist_ListWatchlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchlistServer).ListWatchlist(ctx, req.(*ListWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watchlist_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchlistServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watchlist_AddMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchlistServer).AddMovie(ctx, req.(*AddMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watchlist_RemoveMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchlistServer).RemoveMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watchlist_RemoveMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchlistServer).RemoveMovie(ctx, req.(*RemoveMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Watchlist_ServiceDesc is the grpc.ServiceDesc for Watchlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watchlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "watchlist.Watchlist",
	HandlerType: (*WatchlistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWatchlist",
			Handler:    _Watchlist_ListWatchlist_Handler,
		},
		{
			MethodName: "AddMovie",
			Handler:    _Watchlist_AddMovie_Handler,
		},
		{
			MethodName: "RemoveMovie",
			Handler:    _Watchlist_RemoveMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/watchlist/watchlist.proto",
}
