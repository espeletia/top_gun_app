package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

func (r *tournamentResolver) Owner(ctx context.Context, obj *model.Tournament) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tournamentResolver) Events(ctx context.Context, obj *model.Tournament) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tournament returns generated.TournamentResolver implementation.
func (r *Resolver) Tournament() generated.TournamentResolver { return &tournamentResolver{r} }

type tournamentResolver struct{ *Resolver }
