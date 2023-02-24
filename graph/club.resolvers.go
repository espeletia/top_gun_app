package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

func (r *clubResolver) Owner(ctx context.Context, obj *model.Club) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *clubResolver) Members(ctx context.Context, obj *model.Club) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Club returns generated.ClubResolver implementation.
func (r *Resolver) Club() generated.ClubResolver { return &clubResolver{r} }

type clubResolver struct{ *Resolver }
