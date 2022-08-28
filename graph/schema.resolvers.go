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

func (r *eventResolver) Tournament(ctx context.Context, obj *model.Event) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Referees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) Athletes(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *matchResolver) LeftAthlete(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *matchResolver) RightAthlete(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *matchResolver) Referee(ctx context.Context, obj *model.Match) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTournament(ctx context.Context, input model.CreateTournamentInput) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pooleResolver) Referee(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pooleResolver) Athletes(ctx context.Context, obj *model.Poole) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pooleResolver) Matches(ctx context.Context, obj *model.Poole) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllTournaments(ctx context.Context) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tableauResolver) Matches(ctx context.Context, obj *model.Tableau) ([]*model.Match, error) {
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

// Club returns generated.ClubResolver implementation.
func (r *Resolver) Club() generated.ClubResolver { return &clubResolver{r} }

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Match returns generated.MatchResolver implementation.
func (r *Resolver) Match() generated.MatchResolver { return &matchResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Poole returns generated.PooleResolver implementation.
func (r *Resolver) Poole() generated.PooleResolver { return &pooleResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tableau returns generated.TableauResolver implementation.
func (r *Resolver) Tableau() generated.TableauResolver { return &tableauResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type clubResolver struct{ *Resolver }
type eventResolver struct{ *Resolver }
type matchResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type pooleResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tableauResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
