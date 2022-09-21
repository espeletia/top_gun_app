package graph

import (
	"FenceLive/graph/model"
	"FenceLive/internal/domain"
	"strconv"
	"time"
)

func NewInputMapper() *GqlInputMapper {
	return &GqlInputMapper{}
}

type GqlInputMapper struct {
}

func (gim GqlInputMapper) MapUser(input model.CreateUserInput) (*domain.UserData, error) {
	dateString := input.BornIn
	date, error := time.Parse("2006-01-02", dateString)
	userData := &domain.UserData{
		Email:       input.Email,
		Username:    input.UserName,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Hash:        input.HashedPassword,
		Nationality: input.Nationality,
		BornIn:      date,
	}
	return userData, error
}

func (gim GqlInputMapper) MapTournament(input model.CreateTournamentInput) (*domain.TournamentData, []*domain.EventData, error) {
	ownId, err := strconv.Atoi(input.OwnerID)
	if err != nil {
		return nil, nil, err
	}
	tournmtData := domain.TournamentData{
		Start:       time.Unix(input.Start, 0),
		End:         time.Unix(input.End, 0),
		Name:        input.Name,
		Location:    nil,
		City:        input.City,
		Country:     input.Country,
		OwnerId:     int64(ownId),
		Description: input.Description,
	}
	if input.Location != nil {
		tournmtData.Location = &domain.Location{
			Lon:     input.Location.Lon,
			Lat:     input.Location.Lat,
			Address: input.Location.Address,
		}
	}
	var evntData []*domain.EventData
	for _, evnt := range input.Events {
		event, err := gim.MapEvent(*evnt)
		if err != nil {
			return nil, nil, err
		}

		evntData = append(evntData, event)
	}
	return &tournmtData, evntData, nil
}

func (gim GqlInputMapper) MapEvent(input model.CreateEventInput) (*domain.EventData, error) {
	mappedEvent := &domain.EventData{
		Name:        input.Name,
		Description: input.Description,
		Start:       time.Unix(input.Start, 0),
		End:         time.Unix(input.End, 0),
		Weapon:      string(input.Details.Weapon),
		Type:        string(input.Details.Type),
		Gender:      string(input.Details.Gender),
		Category:    string(input.Details.Category),
	}
	return mappedEvent, nil
}
