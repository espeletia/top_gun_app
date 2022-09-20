package domain

import (
	"time"
)

type Event struct {
	ID     int64
	Status string
	EventData
}

type EventData struct {
	Name         string
	Description  *string
	TournamentId int64
	Start        time.Time
	End          time.Time
	Weapon       string
	Type         string
	Gender       string
	Category     string
}
