package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

func (r *athleteResolver) User(ctx context.Context, obj *model.Athlete) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Tournament(ctx context.Context, obj *model.Event) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Referees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Athletes(ctx context.Context, obj *model.Event) ([]*model.Athlete, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Pooles(ctx context.Context, obj *model.Event) ([]*model.Poole, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Tableaus(ctx context.Context, obj *model.Event) ([]*model.Tableau, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEvent(ctx context.Context, tournamentID string, input model.CreateEventInput) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEvent(ctx context.Context, eventID string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Athlete returns generated.AthleteResolver implementation.
func (r *Resolver) Athlete() generated.AthleteResolver { return &athleteResolver{r} }

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type athleteResolver struct{ *Resolver }
type eventResolver struct{ *Resolver }
