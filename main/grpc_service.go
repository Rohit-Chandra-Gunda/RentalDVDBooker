package main

import (
	grpcservices "RentalDVDBooker/models/grpcservices"
	"RentalDVDBooker/models/protobuffs"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Could not listen to port 5000: ", err)
	}

	grpcServer := grpc.NewServer()

	RegisterGRPCServers(grpcServer)
	serverErr := grpcServer.Serve(lis)

	if serverErr != nil {
		fmt.Println("Failed to serve:", serverErr)
	}
}

func RegisterGRPCServers(gRPCServer *grpc.Server) {
	protobuffs.RegisterBookingServer(gRPCServer, &grpcservices.BookingService{})
	protobuffs.RegisterCancelBookingServer(gRPCServer, &grpcservices.CancelBookingService{})
	protobuffs.RegisterUserBookingsServer(gRPCServer, &grpcservices.UserBookingsService{})
	protobuffs.RegisterDvdBookingsServer(gRPCServer, &grpcservices.DvdBookingsService{})
}
