package users

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/postgres/public/model"
	"FenceLive/internal/ports/database/gen/postgres/public/table"
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
)

func NewUserDatabaseStore(db *sql.DB) *UserDatabaseStore {
	return &UserDatabaseStore{
		DB: db,
	}
}

type UserDatabaseStore struct {
	DB *sql.DB
}

func (udbs UserDatabaseStore) CreateUser(ctx context.Context, user domain.UserData) (*domain.User, error) {

	modelUser := model.Users{
		Email:       user.Email,
		Hash:        user.Hash,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		BornIn:      user.BornIn,
		Nationality: user.Nationality,
	}

	stmt := table.Users.INSERT(table.Users.Email, table.Users.Username, table.Users.FirstName, table.Users.LastName, table.Users.Hash, table.Users.Nationality, table.Users.BornIn).
		MODEL(modelUser).
		RETURNING(table.Users.AllColumns)

	var dest struct {
		model.Users
	}

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID: int64(dest.ID),
		UserData: domain.UserData{
			Email:       dest.Email,
			FirstName:   dest.FirstName,
			LastName:    dest.LastName,
			Username:    dest.Username,
			Hash:        dest.Hash,
			BornIn:      dest.BornIn,
			Nationality: dest.Nationality}}, nil
}

func (udbs UserDatabaseStore) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users)

	var dest []struct {
		model.Users
	}

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for _, user := range dest {
		tempUser := domain.User{
			ID: int64(user.ID),
			UserData: domain.UserData{
				BornIn:      user.BornIn,
				Email:       user.Email,
				Username:    user.Username,
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				Hash:        user.Hash,
				Nationality: user.Nationality,
			},
		}
		users = append(users, &tempUser)
	}
	return users, nil
}

func (udbs UserDatabaseStore) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(table.Users.ID.EQ(postgres.Int(id)))

	var dest []struct {
		model.Users
	}

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) < 1 {
		return nil, domain.UserNotFound
	}

	return &domain.User{
		ID: int64(dest[0].ID),
		UserData: domain.UserData{
			Email:       dest[0].Email,
			FirstName:   dest[0].FirstName,
			LastName:    dest[0].LastName,
			Username:    dest[0].Username,
			Hash:        dest[0].Hash,
			BornIn:      dest[0].BornIn,
			Nationality: dest[0].Nationality}}, nil
}
