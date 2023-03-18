package domain

import "time"

const TournamentStatusCreated = "CREATED"

const TournamentNotFound = Error("Tournament does not exist")

type Tournament struct {
	Id     int64
	Status string
	TournamentData
}

type TournamentData struct {
	Start       time.Time
	End         time.Time
	Name        string
	Location    *Location
	City        string
	Country     string
	OwnerId     int64
	Description *string
}

type Location struct {
	Lat     float64
	Lon     float64
	Address string
}
