package grpcservers

import (
	"RentalDVDBooker/dao"
	"RentalDVDBooker/models/protobuffs"
	"context"
)

type CancelBookingService struct{}

func (s *CancelBookingService) CancelBooking(ctx context.Context,
	request *protobuffs.CancelRequest) (*protobuffs.CancelResponse, error) {
	bookingId := request.GetBookingId()

	success, err := dao.CancelBooking(bookingId)
	response := protobuffs.CancelResponse{}
	response.Success = success

	if err != nil {
		response.ErrorMsg = err.Error()
	}

	return &response, err
}
