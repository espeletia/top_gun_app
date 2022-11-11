package graph

import (
	"FenceLive/graph/model"
	"FenceLive/internal/domain"
	"strconv"
)

func NweGqlMapper() *GqlMapper {
	return &GqlMapper{}
}

type GqlMapper struct {
}

func (gm GqlMapper) MapUser(user *domain.User) (*model.User, error) {
	mappedUser := &model.User{
		ID:        strconv.Itoa(int(user.ID)),
		BornIn:    user.BornIn.String(),
		Email:     user.Email,
		UserName:  user.Username,
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
	}
	return mappedUser, nil
}

func (gm GqlMapper) MapUserArray(users []*domain.User) ([]*model.User, error) {
	var mappedUserArray []*model.User
	for _, user := range users {
		mappedUser, err := gm.MapUser(user)
		if err != nil {
			return nil, err
		}
		mappedUserArray = append(mappedUserArray, mappedUser)
	}
	return mappedUserArray, nil
}

func (gm GqlMapper) MapTournament(tournament *domain.Tournament) (*model.Tournament, error) {
	mappedTournament := &model.Tournament{
		ID:    strconv.Itoa(int(tournament.Id)),
		Start: int64(tournament.Start.Unix()),
		End:   int64(tournament.End.Unix()),
		Name:  tournament.Name,
		Location: &model.Location{
			Lat:     tournament.Location.Lat,
			Lon:     tournament.Location.Lon,
			Address: tournament.Location.Address,
		},
		City:        tournament.City,
		Status:      model.TournamentStatus(tournament.Status),
		Description: tournament.Description,
		OwnerID:     strconv.Itoa(int(tournament.OwnerId)),
	}
	return mappedTournament, nil
}

func (gm GqlMapper) MapTournamentArray(tournament []*domain.Tournament) ([]*model.Tournament, error) {
	var mappedTournamentArray []*model.Tournament
	for _, trn := range tournament {
		mappedTournament, err := gm.MapTournament(trn)
		if err != nil {
			return nil, err
		}
		mappedTournamentArray = append(mappedTournamentArray, mappedTournament)
	}
	return mappedTournamentArray, nil
}

func (gm GqlMapper) MapEvent(event *domain.Event) (*model.Event, error) {
	mappedEvent := &model.Event{
		ID:           strconv.Itoa(int(event.ID)),
		Start:        int64(event.Start.Unix()),
		End:          int64(event.End.Unix()),
		Description:  event.Description,
		TournamentID: strconv.Itoa(int(event.TournamentId)),
		Details: &model.EventDetails{
			Weapon:   model.EventWeapon(event.Weapon),
			Type:     model.EventType(event.Type),
			Gender:   model.EventGenderMix(event.Gender),
			Category: model.EventAgeCategory(event.Category),
		},
	}
	return mappedEvent, nil
}

func (gm GqlMapper) MapEventArray(events []*domain.Event) ([]*model.Event, error) {
	var mappedEventArray []*model.Event
	for _, evnt := range events {
		mappedEvent, err := gm.MapEvent(evnt)
		if err != nil {
			return nil, err
		}
		mappedEventArray = append(mappedEventArray, mappedEvent)
	}
	return mappedEventArray, nil
}
