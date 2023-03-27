package domain

import (
	"time"
)

var (
	UserRoleDefault = "USER"
	UserRoleAdmin   = "ADMIN"
)

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

type LoginCreds struct {
	Email    string
	Password string
}
