package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
)

func NewEventUsecase(tsi database.TournamentStoreInterface) *TournamentUsecase {
	return &TournamentUsecase{
		store: tsi,
	}
}

type EventUsecase struct {
	store database.EventStoreInterface
}

func (eu EventUsecase) CreateEvent(ctx context.Context, event domain.EventData, tournamentId int64) (*domain.Event, error) {
	return eu.store.CreateEvent(ctx, event, tournamentId)
}

func (eu EventUsecase) GetByTournamentId(ctx context.Context, tournamentId int64) ([]*domain.Event, error) {
	return eu.store.GetByTournamentId(ctx, tournamentId)
}
