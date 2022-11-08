package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
)

func NewTournamentUsecase(tsi database.TournamentStoreInterface) TournamentUsecase {
	return TournamentUsecase{
		store: tsi,
	}
}

type TournamentUsecase struct {
	store database.TournamentStoreInterface
}

func (tu TournamentUsecase) CreateTournament(ctx context.Context, tournData domain.TournamentData) (*domain.Tournament, error) {
	return tu.store.CreateTournament(ctx, tournData)
}

func (tu TournamentUsecase) GetByTournamentId(ctx context.Context, id int64) (*domain.Tournament, error) {
	return tu.store.GetByTournamentId(ctx, id)
}

func (tu TournamentUsecase) GetAllTournaments(ctx context.Context) ([]*domain.Tournament, error) {
	return tu.store.GetAllTournaments(ctx)
}

func (tu TournamentUsecase) UpdateTournamentData(ctx context.Context, tournamentId int64, tournamentData domain.TournamentData) (*domain.Tournament, error) {
	return tu.store.UpdateTournamentData(ctx, tournamentId, tournamentData)
}
