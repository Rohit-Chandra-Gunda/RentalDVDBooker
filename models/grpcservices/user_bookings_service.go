package grpcservers

import (
	"RentalDVDBooker/dao"
	"RentalDVDBooker/helpers"
	"RentalDVDBooker/models/protobuffs"
	"context"
)

type UserBookingsService struct{}

func (s *UserBookingsService) GetUserBookings(ctx context.Context,
	request *protobuffs.UserBookingsRequest) (*protobuffs.UserBookingsResponse, error) {
	userId := request.GetUserId()
	status := request.GetStatus()

	userBookings, err := dao.GetUserBookings(userId, status)

	response := protobuffs.UserBookingsResponse{}
	var responseBookings []*protobuffs.Booking

	if err != nil {
		response.ErrorMsg = err.Error()
		responseBookings = make([]*protobuffs.Booking, 0)
	} else {
		responseBookings = helpers.ConvertBookingArrayToProto(userBookings)
	}

	response.Bookings = responseBookings

	return &response, err
}
