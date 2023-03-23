package graph

import (
	"FenceLive/internal/usecases/events"
	tournament "FenceLive/internal/usecases/tournaments"
	"FenceLive/internal/usecases/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users       users.UserUsecase
	Tournaments tournament.TournamentUsecase
	Events      events.EventUsecase

	Mapper      *GqlMapper
	InputMapper *GqlInputMapper
}
