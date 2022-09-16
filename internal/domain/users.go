package domain

import "time"

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
