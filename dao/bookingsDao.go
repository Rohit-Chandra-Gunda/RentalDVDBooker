package dao

import (
	"RentalDVDBooker/models"
	"errors"
	"strconv"
)

var bookingList = []*models.Booking{
	&models.Booking{UserId: 1, DvdId: 1, BookingId: 1, Status: true},
	&models.Booking{UserId: 1, DvdId: 2, BookingId: 2, Status: true},
	&models.Booking{UserId: 2, DvdId: 1, BookingId: 3, Status: true},
	&models.Booking{UserId: 2, DvdId: 2, BookingId: 4, Status: false},
}

var nextBookingId uint64 = 5

func getUserBookings(userId uint64) []*models.Booking {
	userBookings := []*models.Booking{}

	for _, booking := range bookingList {
		if booking.UserId == userId && booking.Status {
			userBookings = append(userBookings, booking)
		}
	}

	return userBookings
}

func getBooking(bookingId uint64) (*models.Booking, error) {
	for _, booking := range bookingList {
		if booking.BookingId == bookingId {
			return booking, nil
		}
	}

	err := errors.New("No booking found with the given id")
	return nil, err
}

func getSpecificBooking(userId, dvdId uint64) (*models.Booking, error) {
	for _, booking := range bookingList {
		if booking.UserId == userId && booking.DvdId == dvdId && booking.Status {
			return booking, nil
		}
	}

	err := errors.New("No booking found with the given id")
	return nil, err
}

func AddBooking(dvdId, userId uint64) (uint64, error) {
	userBookings := getUserBookings(userId)

	for _, booking := range userBookings {
		if booking.DvdId == dvdId {
			err := errors.New("There is an existing booking with userId:" + strconv.FormatUint(userId, 10) + " and dvdId:" + strconv.FormatUint(dvdId, 10))
			return 0, err
		}
	}

	bookingId := nextBookingId
	nextBookingId++
	bookingList = append(bookingList, &models.Booking{UserId: userId, DvdId: dvdId, BookingId: bookingId, Status: true})

	return bookingId, nil
}

func CancelBooking(bookingId uint64) (bool, error) {
	booking, err := getBooking(bookingId)

	if err == nil {
		if !booking.Status {
			err = errors.New("The booking with bookingId:" + strconv.FormatUint(bookingId, 10) + " is already cancelled")
			return false, err
		}

		booking.Status = false
		return true, nil
	} else {
		return false, err
	}
}

func CancelSpecificBooking(userId, dvdId uint64) (bool, error) {
	booking, err := getSpecificBooking(userId, dvdId)

	if err == nil {
		booking.Status = false
		return true, nil
	} else {
		return false, err
	}
}
