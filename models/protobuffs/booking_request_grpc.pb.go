// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: booking_request.proto

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
	Booker_BookDvd_FullMethodName = "/protobuffs.Booker/BookDvd"
)

// BookerClient is the client API for Booker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookerClient interface {
	BookDvd(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
}

type bookerClient struct {
	cc grpc.ClientConnInterface
}

func NewBookerClient(cc grpc.ClientConnInterface) BookerClient {
	return &bookerClient{cc}
}

func (c *bookerClient) BookDvd(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, Booker_BookDvd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookerServer is the server API for Booker service.
// All implementations should embed UnimplementedBookerServer
// for forward compatibility
type BookerServer interface {
	BookDvd(context.Context, *BookingRequest) (*BookingResponse, error)
}

// UnimplementedBookerServer should be embedded to have forward compatible implementations.
type UnimplementedBookerServer struct {
}

func (UnimplementedBookerServer) BookDvd(context.Context, *BookingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookDvd not implemented")
}

// UnsafeBookerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookerServer will
// result in compilation errors.
type UnsafeBookerServer interface {
	mustEmbedUnimplementedBookerServer()
}

func RegisterBookerServer(s grpc.ServiceRegistrar, srv BookerServer) {
	s.RegisterService(&Booker_ServiceDesc, srv)
}

func _Booker_BookDvd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookerServer).BookDvd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booker_BookDvd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookerServer).BookDvd(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Booker_ServiceDesc is the grpc.ServiceDesc for Booker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Booker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuffs.Booker",
	HandlerType: (*BookerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookDvd",
			Handler:    _Booker_BookDvd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking_request.proto",
}
