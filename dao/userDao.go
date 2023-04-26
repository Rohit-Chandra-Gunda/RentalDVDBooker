package dao

import (
	"RentalDVDBooker/models"
	"errors"
	"strconv"
)

var userList = []models.User{
	{UserId: 1, IsActive: true, Name: "user1", City: "city1", State: "state1", Country: "country1"},
	{UserId: 2, IsActive: true, Name: "user2", City: "city2", State: "state2", Country: "country2"},
	{UserId: 3, IsActive: false, Name: "user3", City: "city3", State: "state1", Country: "country1"},
}

func GetUser(userId uint64) (models.User, error) {
	for _, user := range userList {
		if user.UserId == userId {
			return user, nil
		}
	}

	err := errors.New("There is no user with id:" + strconv.FormatUint(userId, 10))
	return models.User{}, err
}

func IsActiveUser(userId uint64) (bool, error) {
	user, err := GetUser(userId)

	if err == nil {
		if user.IsActive {
			return true, nil
		} else {
			err = errors.New("The user with id:" + strconv.FormatUint(userId, 10) + " is not active")
			return false, err
		}
	}

	return false, err
}
