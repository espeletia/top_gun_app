package users

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/fencelive/public/model"
	"FenceLive/internal/ports/database/gen/fencelive/public/table"
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
		Role: &domain.UserRoleDefault,
	}

	stmt := table.Users.INSERT(table.Users.Email, table.Users.Username, table.Users.FirstName, table.Users.LastName, table.Users.Hash, table.Users.Nationality, table.Users.BornIn).
		MODEL(modelUser).
		RETURNING(table.Users.AllColumns)

	var dest model.Users
	

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return mapUserFromDB(dest)
}

func (udbs UserDatabaseStore) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users)

	var dest []model.Users

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for _, user := range dest {
		tempUser, _ := mapUserFromDB(user)
		users = append(users, tempUser)
	}
	return users, nil
}

func (udbs UserDatabaseStore) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(table.Users.ID.EQ(postgres.Int(id)))

	var dest []model.Users

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) < 1 {
		return nil, domain.UserNotFound
	}

	return mapUserFromDB(dest[0])
}

func (udbs UserDatabaseStore) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	stmt := table.Users.SELECT(table.Users.AllColumns).FROM(table.Users).WHERE(table.Users.Email.EQ(postgres.String(email)))

	var dest []model.Users

	err := stmt.Query(udbs.DB, &dest)
	if err != nil {
		return nil, err
	}
	if len(dest) < 1 {
		return nil, domain.UserNotFound
	}

	return mapUserFromDB(dest[0])
}

func mapUserFromDB(usr model.Users) (*domain.User, error) {
	return &domain.User{
		ID: int64(usr.ID),
		UserData: domain.UserData{
			Email:       usr.Email,
			FirstName:   usr.FirstName,
			LastName:    usr.LastName,
			Username:    usr.Username,
			Hash:        usr.Hash,
			BornIn:      usr.BornIn,
			Nationality: usr.Nationality}}, nil
}