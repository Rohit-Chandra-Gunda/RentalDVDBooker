package main

import (
	"RentalDVDBooker/server"
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

	server.RegisterGRPCServers(grpcServer)
	serverErr := grpcServer.Serve(lis)

	if serverErr != nil {
		fmt.Println("Failed to serve:", serverErr)
	}
}
