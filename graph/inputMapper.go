package graph

import (
	"FenceLive/graph/model"
	"FenceLive/internal/domain"
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

func (gim GqlInputMapper) MapTournament(input model.CreateTournamentInput) (*domain.TournamentData, error) {
	return nil, nil
}
