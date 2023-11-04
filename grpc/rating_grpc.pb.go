// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: grpc/rating.proto

package grpc

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

// RatingClient is the client API for Rating service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingClient interface {
	GetRatingByUserId(ctx context.Context, in *GetRatingByUserIdRequest, opts ...grpc.CallOption) (*GetRatingByUserIdResponce, error)
}

type ratingClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingClient(cc grpc.ClientConnInterface) RatingClient {
	return &ratingClient{cc}
}

func (c *ratingClient) GetRatingByUserId(ctx context.Context, in *GetRatingByUserIdRequest, opts ...grpc.CallOption) (*GetRatingByUserIdResponce, error) {
	out := new(GetRatingByUserIdResponce)
	err := c.cc.Invoke(ctx, "/Rating/GetRatingByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServer is the server API for Rating service.
// All implementations must embed UnimplementedRatingServer
// for forward compatibility
type RatingServer interface {
	GetRatingByUserId(context.Context, *GetRatingByUserIdRequest) (*GetRatingByUserIdResponce, error)
	mustEmbedUnimplementedRatingServer()
}

// UnimplementedRatingServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServer struct {
}

func (UnimplementedRatingServer) GetRatingByUserId(context.Context, *GetRatingByUserIdRequest) (*GetRatingByUserIdResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingByUserId not implemented")
}
func (UnimplementedRatingServer) mustEmbedUnimplementedRatingServer() {}

// UnsafeRatingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServer will
// result in compilation errors.
type UnsafeRatingServer interface {
	mustEmbedUnimplementedRatingServer()
}

func RegisterRatingServer(s grpc.ServiceRegistrar, srv RatingServer) {
	s.RegisterService(&Rating_ServiceDesc, srv)
}

func _Rating_GetRatingByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).GetRatingByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Rating/GetRatingByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).GetRatingByUserId(ctx, req.(*GetRatingByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rating_ServiceDesc is the grpc.ServiceDesc for Rating service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rating_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Rating",
	HandlerType: (*RatingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRatingByUserId",
			Handler:    _Rating_GetRatingByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/rating.proto",
}
