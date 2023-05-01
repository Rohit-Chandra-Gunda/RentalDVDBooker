package dao

import (
	"RentalDVDBooker/models"
	"errors"
	"strconv"
)

var dvdList = []models.Dvd{
	{DvdId: 1, IsAvailable: true, Name: "dvd1"},
	{DvdId: 2, IsAvailable: true, Name: "dvd2"},
	{DvdId: 3, IsAvailable: false, Name: "dvd3"},
}

func GetDvd(dvdId uint64) (models.Dvd, error) {
	for _, dvd := range dvdList {
		if dvd.DvdId == dvdId {
			return dvd, nil
		}
	}

	err := errors.New("There is no dvd with id:" + strconv.FormatUint(dvdId, 10))
	return models.Dvd{}, err
}

func IsDvdAvailable(dvdId uint64) (bool, error) {
	dvd, err := GetDvd(dvdId)

	if err == nil {
		if dvd.IsAvailable {
			return true, nil
		} else {
			err = errors.New("The dvd with id:" + strconv.FormatUint(dvdId, 10) + " is not available")
			return false, err
		}
	}

	return false, err
}
