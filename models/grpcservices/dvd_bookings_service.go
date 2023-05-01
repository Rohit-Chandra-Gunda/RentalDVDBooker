package grpcservers

import (
	"RentalDVDBooker/dao"
	"RentalDVDBooker/helpers"
	"RentalDVDBooker/models/protobuffs"
	"context"
)

type DvdBookingsService struct{}

func (s *DvdBookingsService) GetDvdBookings(ctx context.Context,
	request *protobuffs.DvdBookingsRequest) (*protobuffs.DvdBookingsResponse, error) {
	dvdId := request.GetDvdId()
	status := request.GetStatus()

	dvdBookings, err := dao.GetDvdBookings(dvdId, status)

	response := protobuffs.DvdBookingsResponse{}
	var responseBookings []*protobuffs.Booking

	if err != nil {
		response.ErrorMsg = err.Error()
		responseBookings = make([]*protobuffs.Booking, 0)
	} else {
		responseBookings = helpers.ConvertBookingArrayToProto(dvdBookings)
	}

	response.Bookings = responseBookings

	return &response, err
}
