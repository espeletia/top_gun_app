package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

// LeftAthlete is the resolver for the LeftAthlete field.
func (r *matchResolver) LeftAthlete(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// RightAthlete is the resolver for the RightAthlete field.
func (r *matchResolver) RightAthlete(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Referee is the resolver for the Referee field.
func (r *matchResolver) Referee(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Match returns generated.MatchResolver implementation.
func (r *Resolver) Match() generated.MatchResolver { return &matchResolver{r} }

type matchResolver struct{ *Resolver }
