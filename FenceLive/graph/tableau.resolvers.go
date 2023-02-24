package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

// Matches is the resolver for the Matches field.
func (r *tableauResolver) Matches(ctx context.Context, obj *model.Tableau) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tableau returns generated.TableauResolver implementation.
func (r *Resolver) Tableau() generated.TableauResolver { return &tableauResolver{r} }

type tableauResolver struct{ *Resolver }