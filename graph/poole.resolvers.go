package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

func (r *pooleResolver) Referee(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pooleResolver) Athletes(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pooleResolver) Matches(ctx context.Context, obj *model.Poole) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

// Poole returns generated.PooleResolver implementation.
func (r *Resolver) Poole() generated.PooleResolver { return &pooleResolver{r} }

type pooleResolver struct{ *Resolver }
