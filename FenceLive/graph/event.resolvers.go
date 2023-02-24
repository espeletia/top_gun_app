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

// User is the resolver for the User field.
func (r *athleteResolver) User(ctx context.Context, obj *model.Athlete) (*model.User, error) {
	userId, err := strconv.Atoi(obj.UserID)
	if err != nil {
		return nil, err
	}
	user, err := r.Users.GetUserById(ctx, int64(userId))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapUser(user)
}

// Tournament is the resolver for the Tournament field.
func (r *eventResolver) Tournament(ctx context.Context, obj *model.Event) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

// Referees is the resolver for the Referees field.
func (r *eventResolver) Referees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Athletes is the resolver for the Athletes field.
func (r *eventResolver) Athletes(ctx context.Context, obj *model.Event) ([]*model.Athlete, error) {
	eventId, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}
	athletes, err := r.Events.GetAllAthletes(ctx, int64(eventId))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapAthleteArray(athletes)
}

// Pooles is the resolver for the Pooles field.
func (r *eventResolver) Pooles(ctx context.Context, obj *model.Event) ([]*model.Poole, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tableaus is the resolver for the Tableaus field.
func (r *eventResolver) Tableaus(ctx context.Context, obj *model.Event) ([]*model.Tableau, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateEvent is the resolver for the CreateEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, tournamentID string, input model.CreateEventInput) (*model.Event, error) {
	event, err := r.InputMapper.MapEvent(input)
	if err != nil {
		return nil, err
	}
	tournId, err := strconv.Atoi(tournamentID)
	if err != nil {
		return nil, err
	}
	storedEvent, err := r.Events.CreateEvent(ctx, *event, int64(tournId))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapEvent(storedEvent)
}

// GetEvent is the resolver for the GetEvent field.
func (r *queryResolver) GetEvent(ctx context.Context, eventID string) (*model.Event, error) {
	eventId, err := strconv.Atoi(eventID)
	if err != nil {
		return nil, err
	}
	event, err := r.Events.GetEventById(ctx, int64(eventId))
	if err != nil {
		return nil, err
	}
	return r.Mapper.MapEvent(event)
}

// Athlete returns generated.AthleteResolver implementation.
func (r *Resolver) Athlete() generated.AthleteResolver { return &athleteResolver{r} }

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type athleteResolver struct{ *Resolver }
type eventResolver struct{ *Resolver }
