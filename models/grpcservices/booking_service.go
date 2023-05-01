package grpcservers

import (
	"RentalDVDBooker/dao"
	"RentalDVDBooker/models/protobuffs"
	"context"
)

type BookingService struct{}

func (s *BookingService) BookDvd(ctx context.Context,
	request *protobuffs.BookingRequest) (*protobuffs.BookingResponse, error) {
	userId := request.GetUserId()
	dvdId := request.GetDvdId()

	bookingId, err := dao.AddBooking(userId, dvdId)
	response := protobuffs.BookingResponse{}

	if err != nil {
		response.ErrorMsg = err.Error()
		response.Success = false
	} else {
		response.Success = true
		response.BookingId = bookingId
	}

	return &response, err
}
