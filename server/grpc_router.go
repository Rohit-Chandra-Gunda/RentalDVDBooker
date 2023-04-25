package server

import (
	"RentalDVDBooker/models/protobuffs"
	"context"

	"google.golang.org/grpc"
)

func RegisterGRPCServers(gRPCServer *grpc.Server) {
	protobuffs.RegisterBookerServer(gRPCServer, &BookerServer{})
	protobuffs.RegisterCancelBookingServer(gRPCServer, &CancelBookingServer{})
}

type BookerServer struct{}

func (s *BookerServer) BookDvd(ctx context.Context,
	request *protobuffs.BookingRequest) (*protobuffs.BookingResponse, error) {

}

type CancelBookingServer struct{}

func (s *CancelBookingServer) CancelBooking(ctx context.Context,
	request protobuffs.CancelRequest) (*protobuffs.CancelResponse, error) {

}
