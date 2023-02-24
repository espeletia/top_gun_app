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

func (gim GqlInputMapper) MapUser(input model.CreateUserInput) domain.UserData {
	dateString := input.BornIn
	date, _ := time.Parse("2006-01-02", dateString)
	userData := domain.UserData{
		Email:       input.Email,
		Username:    input.UserName,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Hash:        input.HashedPassword,
		Nationality: input.Nationality,
		BornIn:      date,
	}
	return userData
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
	athletes, err := gim.MapEventUserArray(input.Athletes)
	if err != nil {
		return nil, err
	}
	mappedEvent := &domain.EventData{
		Name:        input.Name,
		Description: input.Description,
		Start:       time.Unix(input.Start, 0),
		End:         time.Unix(input.End, 0),
		Weapon:      string(input.Details.Weapon),
		Type:        string(input.Details.Type),
		Gender:      string(input.Details.Gender),
		Category:    string(input.Details.Category),
		Athletes:    athletes,
	}
	return mappedEvent, nil
}

func (gim GqlInputMapper) MapTournamentUpdate(input model.UpdateTournamentInput) (*domain.TournamentData, error) {
	ownId, err := strconv.Atoi(input.OwnerID)
	if err != nil {
		return nil, err
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
	return &tournmtData, nil
}

func (gim GqlInputMapper) MapEventUser(input model.AthleteSeedingInput) (*domain.Athlete, error) {
	userId, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.Athlete{
		UserID:       int64(userId),
		PooleSeeding: input.Seed,
	}, nil
}

func (gim GqlInputMapper) MapEventUserArray(input []*model.AthleteSeedingInput) ([]*domain.Athlete, error) {
	var mappedAthletes []*domain.Athlete
	for _, athlete := range input {
		mappedAthlete, err := gim.MapEventUser(*athlete)
		if err != nil {
			return nil, err
		}
		mappedAthletes = append(mappedAthletes, mappedAthlete)
	}
	return mappedAthletes, nil
}
