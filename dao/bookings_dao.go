package dao

import (
	"RentalDVDBooker/models"
	"errors"
	"strconv"
)

var bookingList = []*models.Booking{
	{UserId: 1, DvdId: 1, BookingId: 1, Status: true},
	{UserId: 1, DvdId: 2, BookingId: 2, Status: true},
	{UserId: 2, DvdId: 1, BookingId: 3, Status: true},
	{UserId: 2, DvdId: 2, BookingId: 4, Status: false},
}

var nextBookingId uint64 = 5

func GetUserBookings(userId uint64, status bool) ([]models.Booking, error) {
	if isActive, err := IsActiveUser(userId); !isActive {
		return []models.Booking{}, err
	}

	userBookings := []models.Booking{}

	for _, booking := range bookingList {
		if booking.UserId == userId && booking.Status == status {
			userBookings = append(userBookings, *booking)
		}
	}

	return userBookings, nil
}

func GetDvdBookings(dvdId uint64, status bool) ([]models.Booking, error) {
	if isActive, err := IsActiveUser(dvdId); !isActive {
		return []models.Booking{}, err
	}

	dvdBookings := []models.Booking{}

	for _, booking := range bookingList {
		if booking.DvdId == dvdId && booking.Status == status {
			dvdBookings = append(dvdBookings, *booking)
		}
	}

	return dvdBookings, nil
}

func GetBookingFromId(bookingId uint64) (*models.Booking, error) {
	for _, booking := range bookingList {
		if booking.BookingId == bookingId {
			return booking, nil
		}
	}

	err := errors.New("No booking found with the id:" + strconv.FormatUint(bookingId, 10))
	return nil, err
}

func GetBooking(userId, dvdId uint64) (*models.Booking, error) {
	for _, booking := range bookingList {
		if booking.UserId == userId && booking.DvdId == dvdId && booking.Status {
			return booking, nil
		}
	}

	err := errors.New("No active booking found with the user id:" + strconv.FormatUint(userId, 10) + " and dvd id:" + strconv.FormatUint(userId, 10))
	return nil, err
}

func AddBooking(userId, dvdId uint64) (uint64, error) {
	if isActive, err := IsActiveUser(userId); !isActive {
		return 0, err
	}

	if isAvailable, err := IsDvdAvailable(dvdId); !isAvailable {
		return 0, err
	}

	booking, err := GetBooking(userId, dvdId)

	if err == nil {
		err := errors.New("There is an existing booking with user id:" + strconv.FormatUint(userId, 10) + " and dvd id:" + strconv.FormatUint(dvdId, 10))
		return booking.BookingId, err
	}

	bookingId := nextBookingId
	nextBookingId++
	bookingList = append(bookingList, &models.Booking{UserId: userId, DvdId: dvdId, BookingId: bookingId, Status: true})

	return bookingId, nil
}

func CancelBooking(bookingId uint64) (bool, error) {
	booking, err := GetBookingFromId(bookingId)

	if err != nil {
		return false, nil
	}

	if !booking.Status {
		err = errors.New("The booking with id:" + strconv.FormatUint(booking.BookingId, 10) + " is not active")
		return false, nil
	}

	booking.Status = false
	return true, nil
}
