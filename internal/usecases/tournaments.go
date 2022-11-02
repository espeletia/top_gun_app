package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
	"fmt"
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
	fmt.Printf("UHHHHHHHHHH\n")
	return tu.store.CreateTournament(ctx, tournData)
} //ok
