package models

type User struct {
	UserId   uint64
	IsActive bool
	Name     string
	City     string
	State    string
	Country  string
}
