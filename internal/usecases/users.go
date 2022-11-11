package usecases

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
)

func NewUserUsecase(usi database.UserStoreInterface) UserUsecase {
	return UserUsecase{
		store: usi,
	}
}

type UserUsecase struct {
	store database.UserStoreInterface
}

func (uu UserUsecase) CreateUser(ctx context.Context, user domain.UserData) (*domain.User, error) {
	stored, err := uu.store.CreateUser(ctx, user)
	return stored, err
}

func (uu UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := uu.store.GetAllUsers(ctx)
	return users, err
}

func (uu UserUsecase) GetUserById(ctx context.Context, Id int64) (*domain.User, error) {
	return uu.store.GetUserById(ctx, Id)
}
