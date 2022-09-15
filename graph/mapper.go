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
