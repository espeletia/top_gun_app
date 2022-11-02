package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
	"fmt"
)

func NewEventUsecase(tsi database.EventStoreInterface) EventUsecase {
	return EventUsecase{
		store: tsi,
	}
}

type EventUsecase struct {
	store database.EventStoreInterface
}

func (eu EventUsecase) CreateEvent(ctx context.Context, event domain.EventData, tournamentId int64) (*domain.Event, error) {
	fmt.Printf("\nhuh\n")
	return eu.store.CreateEvent(ctx, event, tournamentId)
}

func (eu EventUsecase) GetByTournamentId(ctx context.Context, tournamentId int64) ([]*domain.Event, error) {
	return eu.store.GetByTournamentId(ctx, tournamentId)
}
