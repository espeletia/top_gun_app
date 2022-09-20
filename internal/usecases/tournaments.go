package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
)

func NewTournamentUsecase(tsi database.TournamentStoreInterface) *TournamentUsecase {
	return &TournamentUsecase{
		store: tsi,
	}
}

type TournamentUsecase struct {
	store database.TournamentStoreInterface
}

func (tu TournamentUsecase) CreateTournament(ctx context.Context, tournData domain.TournamentData) (*domain.Tournament, error) {
	return tu.store.CreateTournament(ctx, tournData)
} //don't reveal too early and too inconviniently
