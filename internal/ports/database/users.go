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
	stmt := table.Users.INSERT(table.Users.Email, table.Users.Username, table.Users.FirstName, table.Users.LastName, table.Users.Hash).
		VALUES(user.Email, user.Username, user.FirstName, user.LastName, user.Hash).
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
			Email:     dest.Email,
			FirstName: dest.FirstName,
			LastName:  dest.LastName,
			Username:  dest.Username,
			Hash:      dest.Hash,
			BornIn:    dest.BornIn}}, nil
}
