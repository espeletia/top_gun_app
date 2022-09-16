package database

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/FenceLive/public/model"
	"FenceLive/internal/ports/database/gen/FenceLive/public/table"
	"context"
	"database/sql"
)

type UserStoreInterface interface {
	CreateUser(ctx context.Context, user *domain.UserData) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	//GetUser(id string) (*User, error)
	//GetUserByEmail(email string) (*User, error)
	//GetUserByUsername(username string) (*User, error)
	//GetUsers() ([]*User, error)
	//UpdateUser(user *User) error
	//DeleteUser(id string) error
}

func NewUserDatabaseStore(db *sql.DB) *UserDatabaseStore {
	return &UserDatabaseStore{
		db: db,
	}
}

type UserDatabaseStore struct {
	db *sql.DB
}

func (udbs *UserDatabaseStore) CreateUser(ctx context.Context, user *domain.UserData) (*domain.User, error) {

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

	err := stmt.Query(udbs.db, &dest)
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

func (udbs *UserDatabaseStore) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users)

	var dest []struct {
		model.Users
	}

	err := stmt.Query(udbs.db, &dest)
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for _, user := range dest {
		tempUser := *&domain.User{
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
