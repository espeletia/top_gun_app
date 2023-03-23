package domain

import "net/http"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string { return e.Message }

var (
	UserNotFound = Error{Message: "User not found", Code: http.StatusNotFound}
	InvalidPassword = Error{Message: "Password can not be validated", Code: http.StatusNotAcceptable}
	InvalidCredentials = Error{Message: "Invalid credentials", Code: http.StatusUnauthorized}

	TournamentNotFound = Error{Message: "Tournament not found", Code: http.StatusNotFound}

	EventNotFound = Error{Message: "Event not found", Code: http.StatusNotFound}
)
	