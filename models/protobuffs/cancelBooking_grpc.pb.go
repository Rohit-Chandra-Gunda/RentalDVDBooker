// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: cancelBooking.proto

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
	CancelBooking_CancelBooking_FullMethodName = "/protobuffs.CancelBooking/CancelBooking"
)

// CancelBookingClient is the client API for CancelBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CancelBookingClient interface {
	CancelBooking(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type cancelBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewCancelBookingClient(cc grpc.ClientConnInterface) CancelBookingClient {
	return &cancelBookingClient{cc}
}

func (c *cancelBookingClient) CancelBooking(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, CancelBooking_CancelBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CancelBookingServer is the server API for CancelBooking service.
// All implementations must embed UnimplementedCancelBookingServer
// for forward compatibility
type CancelBookingServer interface {
	CancelBooking(context.Context, *CancelRequest) (*CancelResponse, error)
}

// UnimplementedCancelBookingServer must be embedded to have forward compatible implementations.
type UnimplementedCancelBookingServer struct {
}

func (UnimplementedCancelBookingServer) CancelBooking(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}

// UnsafeCancelBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CancelBookingServer will
// result in compilation errors.
type UnsafeCancelBookingServer interface {
	mustEmbedUnimplementedCancelBookingServer()
}

func RegisterCancelBookingServer(s grpc.ServiceRegistrar, srv CancelBookingServer) {
	s.RegisterService(&CancelBooking_ServiceDesc, srv)
}

func _CancelBooking_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CancelBookingServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CancelBooking_CancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CancelBookingServer).CancelBooking(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CancelBooking_ServiceDesc is the grpc.ServiceDesc for CancelBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CancelBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuffs.CancelBooking",
	HandlerType: (*CancelBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelBooking",
			Handler:    _CancelBooking_CancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cancelBooking.proto",
}
