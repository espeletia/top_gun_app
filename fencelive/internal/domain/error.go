package domain

import "net/http"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string { return e.Message }

var (
	InvalidPassword = Error{Message: "Password can not be validated", Code: http.StatusNotAcceptable}
	InvalidCredentials = Error{Message: "Invalid credentials", Code: http.StatusUnauthorized}
	Unauthorized = Error{Message: "Unauthorized", Code: http.StatusUnauthorized}

	UserNotFound = Error{Message: "User not found", Code: http.StatusNotFound}
	TournamentNotFound = Error{Message: "Tournament not found", Code: http.StatusNotFound}

	EventNotFound = Error{Message: "Event not found", Code: http.StatusNotFound}
)
	