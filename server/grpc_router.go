package server

import (
	"RentalDVDBooker/models/protobuffs"

	"google.golang.org/grpc"
)

func RegisterGRPCServers(gRPCServer *grpc.Server) {
	protobuffs.RegisterBookerServer(gRPCServer, &BookerServer)
	protobuffs.RegisterCancelBookingServer(gRPCServer, &CancelBookingServer)
}

type BookerServer struct{}
