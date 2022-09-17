package domain

import "time"

type Error string

func (e Error) Error() string { return string(e) }

const UserNotFound = Error("User does not exist")

type User struct {
	ID int64
	UserData
}

type UserData struct {
	BornIn      time.Time
	Email       string
	Username    string
	FirstName   string
	LastName    string
	Hash        string
	Nationality string
}
