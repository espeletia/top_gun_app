package graph

import "FenceLive/internal/usecases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users       *usecases.UserUsecase
	Tournaments *usecases.TournamentUsecase
	Events      *usecases.EventUsecase

	Mapper      *GqlMapper
	InputMapper *GqlInputMapper
}
