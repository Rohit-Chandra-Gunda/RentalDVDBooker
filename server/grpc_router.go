package server

import (
	"RentalDVDBooker/dao"
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
	userId := request.GetBasicRequest().GetUserId()
	dvdId := request.GetBasicRequest().GetDvdId()

	bookingId, err := dao.AddBooking(userId, dvdId)
	response := protobuffs.BookingResponse{}
	basicResponse := protobuffs.BasicResponse{}

	if err != nil {
		basicResponse.ErrorMsg = err.Error()
		basicResponse.Success = false
	} else {
		basicResponse.Success = true
		basicResponse.BookingId = bookingId
	}

	response.BasicResponse = &basicResponse
	return &response, err
}

type CancelBookingServer struct{}

func (s *CancelBookingServer) CancelBooking(ctx context.Context,
	request *protobuffs.CancelRequest) (*protobuffs.CancelResponse, error) {
	userId := request.GetBasicRequest().GetUserId()
	dvdId := request.GetBasicRequest().GetDvdId()

	bookingId, err := dao.CancelBooking(userId, dvdId)
	response := protobuffs.CancelResponse{}
	basicResponse := protobuffs.BasicResponse{}

	if err != nil {
		basicResponse.ErrorMsg = err.Error()
		basicResponse.Success = false
	} else {
		basicResponse.Success = true
		basicResponse.BookingId = bookingId
	}

	response.BasicResponse = &basicResponse
	return &response, err
}
