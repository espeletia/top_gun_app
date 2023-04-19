package database

import (
	"FenceLive/internal/domain"
	"context"
)

type EventStoreInterface interface {
	CreateEvent(ctx context.Context, event domain.EventData, tournamentId int64) (*domain.Event, error)
	GetByTournamentId(ctx context.Context, tournamentId int64) ([]*domain.Event, error)
	GetAllAthletes(ctx context.Context, eventId int64) ([]*domain.Athlete, error)
	GetEventById(ctx context.Context, eventId int64) (*domain.Event, error)
}

type TournamentStoreInterface interface {
	ListAllTournaments(ctx context.Context, limit int64, offset int64) ([]*domain.Tournament, error)
	CreateTournament(ctx context.Context, TournData domain.TournamentData) (*domain.Tournament, error)
	GetTournamentById(ctx context.Context, id int64) (*domain.Tournament, error)
	GetAllTournaments(ctx context.Context) ([]*domain.Tournament, error)
	UpdateTournamentData(ctx context.Context, tournamentId int64, tournamentData domain.TournamentData) (*domain.Tournament, error)
}

// TODO finish this
type UserStoreInterface interface {
	CreateUser(ctx context.Context, user domain.UserData) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	GetUserById(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	//GetUserByUsername(username string) (*User, error)
	//GetUsers() ([]*User, error)
	//UpdateUser(user *User) error
	//DeleteUser(id string) error
}
