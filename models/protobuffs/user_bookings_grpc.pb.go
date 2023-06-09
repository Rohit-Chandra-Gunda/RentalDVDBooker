// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: user_bookings.proto

package protobuffs

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
	UserBookings_GetUserBookings_FullMethodName = "/protobuffs.UserBookings/GetUserBookings"
)

// UserBookingsClient is the client API for UserBookings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserBookingsClient interface {
	GetUserBookings(ctx context.Context, in *UserBookingsRequest, opts ...grpc.CallOption) (*UserBookingsResponse, error)
}

type userBookingsClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBookingsClient(cc grpc.ClientConnInterface) UserBookingsClient {
	return &userBookingsClient{cc}
}

func (c *userBookingsClient) GetUserBookings(ctx context.Context, in *UserBookingsRequest, opts ...grpc.CallOption) (*UserBookingsResponse, error) {
	out := new(UserBookingsResponse)
	err := c.cc.Invoke(ctx, UserBookings_GetUserBookings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBookingsServer is the server API for UserBookings service.
// All implementations should embed UnimplementedUserBookingsServer
// for forward compatibility
type UserBookingsServer interface {
	GetUserBookings(context.Context, *UserBookingsRequest) (*UserBookingsResponse, error)
}

// UnimplementedUserBookingsServer should be embedded to have forward compatible implementations.
type UnimplementedUserBookingsServer struct {
}

func (UnimplementedUserBookingsServer) GetUserBookings(context.Context, *UserBookingsRequest) (*UserBookingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBookings not implemented")
}

// UnsafeUserBookingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserBookingsServer will
// result in compilation errors.
type UnsafeUserBookingsServer interface {
	mustEmbedUnimplementedUserBookingsServer()
}

func RegisterUserBookingsServer(s grpc.ServiceRegistrar, srv UserBookingsServer) {
	s.RegisterService(&UserBookings_ServiceDesc, srv)
}

func _UserBookings_GetUserBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBookingsServer).GetUserBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBookings_GetUserBookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBookingsServer).GetUserBookings(ctx, req.(*UserBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserBookings_ServiceDesc is the grpc.ServiceDesc for UserBookings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserBookings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuffs.UserBookings",
	HandlerType: (*UserBookingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserBookings",
			Handler:    _UserBookings_GetUserBookings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_bookings.proto",
}
