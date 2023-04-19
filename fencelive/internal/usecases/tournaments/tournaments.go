package tournament

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
	"encoding/base64"
	"encoding/json"
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

func (tu TournamentUsecase) GetTournamentById(ctx context.Context, id int64) (*domain.Tournament, error) {
	return tu.store.GetTournamentById(ctx, id)
}

func (tu TournamentUsecase) GetAllTournaments(ctx context.Context) ([]*domain.Tournament, error) {
	return tu.store.GetAllTournaments(ctx)
}

func (tu TournamentUsecase) UpdateTournamentData(ctx context.Context, tournamentId int64, tournamentData domain.TournamentData) (*domain.Tournament, error) {
	return tu.store.UpdateTournamentData(ctx, tournamentId, tournamentData)
}

func (tu TournamentUsecase) ListAllTournaments(ctx context.Context, limit int64, NextToken string) ([]*domain.Tournament, error) {
	/*
		TODO
		NextToken is a base64 encoded json string that contains the offset for the next query
		Do it this way

		To activate a tournament, we can give control over it to the user or use a watchdog script that will check the start time of the tournament and activate it when the time is right
		Good idea, but we need to be careful with the timezones
		We can use the timezones of the users to calculate the timezones of the tournaments

		thanks copilot

		Good luck adding all this
	*/
	if NextToken == "" {
		return tu.store.ListAllTournaments(ctx, limit, 0)
	}
	//next token is encoded in base64, we need to decode it
	offset, err := base64.StdEncoding.DecodeString(NextToken)
	if err != nil {
		return nil, err
	}
	//next token is a json once decoded, we need to unmarshal it
	NextItem := domain.NextToken{}
	err = json.Unmarshal(offset, &NextItem)
	if err != nil {
		return nil, err
	}

	return tu.store.ListAllTournaments(ctx, limit, NextItem.Offset)
}
