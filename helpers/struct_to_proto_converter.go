package helpers

import (
	"RentalDVDBooker/models"
	"RentalDVDBooker/models/protobuffs"
)

func ConvertBookingToProto(booking models.Booking) *protobuffs.Booking {
	protoBooking := &protobuffs.Booking{}

	protoBooking.UserId = booking.UserId
	protoBooking.DvdId = booking.DvdId
	protoBooking.BookingId = booking.BookingId

	return protoBooking
}

func ConvertBookingArrayToProto(bookings []models.Booking) []*protobuffs.Booking {
	protoBookings := make([]*protobuffs.Booking, len(bookings))

	for index, booking := range bookings {
		protoBookings[index] = ConvertBookingToProto(booking)
	}

	return protoBookings
}
