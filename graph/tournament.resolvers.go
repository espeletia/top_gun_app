package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"FenceLive/internal/domain"
	"context"
	"fmt"
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
	var events []*domain.Event
	for _, event := range eventInput {
		event, err := r.Events.CreateEvent(ctx, *event, tournament.Id)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return 
}

func (r *queryResolver) GetAllTournaments(ctx context.Context) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tournamentResolver) Owner(ctx context.Context, obj *model.Tournament) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tournamentResolver) Events(ctx context.Context, obj *model.Tournament) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tournament returns generated.TournamentResolver implementation.
func (r *Resolver) Tournament() generated.TournamentResolver { return &tournamentResolver{r} }

type tournamentResolver struct{ *Resolver }
