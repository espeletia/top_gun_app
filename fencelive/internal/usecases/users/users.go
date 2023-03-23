package users

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"FenceLive/internal/usecases"
	"FenceLive/internal/usecases/hash"
	"context"
)

func NewUserUsecase(usi database.UserStoreInterface, hash hash.HashUsecase) UserUsecase {
	return UserUsecase{
		store: usi,
		hash:  hash,
	}
}

type UserUsecase struct {
	store database.UserStoreInterface
	hash  usecases.HashUsecaseInterface
}

func (uu UserUsecase) CreateUser(ctx context.Context, user domain.UserData, password string) (*domain.User, error) {
	hashedPassword := uu.hash.HashPassword(password)
	user.Hash = hashedPassword
	stored, err := uu.store.CreateUser(ctx, user)
	return stored, err
}

func (uu UserUsecase) Login(ctx context.Context, creds domain.LoginCreds) (string, error) {
	usr, err := uu.GetUserByEmail(ctx, creds.Email)
	if err != nil {
		return "failure", err
	}
	hashedPassword := uu.hash.HashPassword(creds.Password)
	if hashedPassword != usr.Hash{
		return "You suck", domain.InvalidCredentials
	}
	return "Success", nil
}

func (uu UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := uu.store.GetAllUsers(ctx)
	return users, err
}

func (uu UserUsecase) GetUserById(ctx context.Context, Id int64) (*domain.User, error) {
	return uu.store.GetUserById(ctx, Id)
}

func (uu UserUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uu.store.GetUserByEmail(ctx, email)
}
