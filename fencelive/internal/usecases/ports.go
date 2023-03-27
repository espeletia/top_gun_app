package usecases

import (
	"FenceLive/internal/domain"
	"context"
)

type UserUsecaseInterface interface {
	HashPassword(password string) string
	GetUserById(ctx context.Context, Id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}
