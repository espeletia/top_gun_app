package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"strconv"
)

func (r *mutationResolver) CreateTournament(ctx context.Context, input model.CreateTournamentInput) (*model.Tournament, error) {
	tournamentInput, eventInput, err := r.InputMapper.MapTournament(input)
	if err != nil {
		return nil, err
	}
	tournament, err := r.Tournaments.CreateTournament(ctx, *tournamentInput)
	if err != nil {
		return nil, err
	}
	for _, event := range eventInput {
		_, err := r.Events.CreateEvent(ctx, *event, tournament.Id)
		if err != nil {
			return nil, err
		}
	}
	return r.Mapper.MapTournament(tournament)
}

func (r *mutationResolver) UpdateTournament(ctx context.Context, id string, input model.UpdateTournamentInput) (*model.Tournament, error) {
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	tournamentUpdateMapped, err := r.InputMapper.MapTournamentUpdate(input)
	if err != nil {
		return nil, err
	}
	tournament, err := r.Tournaments.UpdateTournamentData(ctx, int64(tournamentId), *tournamentUpdateMapped)
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapTournament(tournament)
}

func (r *queryResolver) GetAllTournaments(ctx context.Context) ([]*model.Tournament, error) {
	tournaments, err := r.Tournaments.GetAllTournaments(ctx)
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapTournamentArray(tournaments)
}

func (r *queryResolver) GetTournamentByID(ctx context.Context, id string) (*model.Tournament, error) {
	tournamentId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	tournament, err := r.Tournaments.GetByTournamentId(ctx, int64(tournamentId))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapTournament(tournament)
}

func (r *tournamentResolver) Owner(ctx context.Context, obj *model.Tournament) (*model.User, error) {
	ownerID, err := strconv.Atoi(obj.OwnerID)
	if err != nil {
		return nil, err
	}
	owner, err := r.Users.GetUserById(ctx, int64(ownerID))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapUser(owner)
}

func (r *tournamentResolver) Events(ctx context.Context, obj *model.Tournament) ([]*model.Event, error) {
	tournamentId, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}
	events, err := r.Resolver.Events.GetByTournamentId(ctx, int64(tournamentId))
	if err != nil {
		return nil, err
	}
	return r.Resolver.Mapper.MapEventArray(events)
}

// Tournament returns generated.TournamentResolver implementation.
func (r *Resolver) Tournament() generated.TournamentResolver { return &tournamentResolver{r} }

type tournamentResolver struct{ *Resolver }
