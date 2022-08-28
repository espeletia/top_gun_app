package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ParticipatingTournaments(ctx context.Context, obj *model.User) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LikedTournaments(ctx context.Context, obj *model.User) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Following(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Followers(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
