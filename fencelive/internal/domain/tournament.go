package domain

import "time"

const (
	TournamentStatusCreated = "CREATED"
	TournementStatusStarted = "STARTED"
	TournementStatusEnded   = "ENDED"
)

type Tournament struct {
	Id     int64
	Status string
	TournamentData
}

type TournamentFilter struct {
	Start   time.Time
	End     time.Time
	Name    string
	City    string
	Country string
	Status  string
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
