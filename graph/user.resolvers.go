package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"FenceLive/graph/generated"
	"FenceLive/graph/model"
	"context"
	"fmt"
	"strconv"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	userData := r.InputMapper.MapUser(input)
	user, err := r.Users.CreateUser(ctx, userData)
	if err != nil {
		return nil, err
	}
	mappedUser, err := r.Mapper.MapUser(user)
	if err != nil {
		return nil, err
	}
	return mappedUser, nil
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.Users.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapUserArray(users)
}

// GetUserByID is the resolver for the getUserByID field.
func (r *queryResolver) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	Id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	user, err := r.Users.GetUserById(ctx, int64(Id))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapUser(user)
}

// ParticipatingTournaments is the resolver for the ParticipatingTournaments field.
func (r *userResolver) ParticipatingTournaments(ctx context.Context, obj *model.User) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

// LikedTournaments is the resolver for the LikedTournaments field.
func (r *userResolver) LikedTournaments(ctx context.Context, obj *model.User) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

// Following is the resolver for the Following field.
func (r *userResolver) Following(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Followers is the resolver for the Followers field.
func (r *userResolver) Followers(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
