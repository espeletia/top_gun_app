package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

// Referee is the resolver for the Referee field.
func (r *pooleResolver) Referee(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Athletes is the resolver for the Athletes field.
func (r *pooleResolver) Athletes(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Matches is the resolver for the Matches field.
func (r *pooleResolver) Matches(ctx context.Context, obj *model.Poole) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

// Poole returns generated.PooleResolver implementation.
func (r *Resolver) Poole() generated.PooleResolver { return &pooleResolver{r} }

type pooleResolver struct{ *Resolver }