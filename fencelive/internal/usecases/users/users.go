package users

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database"
	"context"
	"crypto/sha512"
	"encoding/hex"

	"go.uber.org/zap"
)

func NewUserUsecase(usi database.UserStoreInterface, salt string) UserUsecase {
	return UserUsecase{
		store: usi,
		salt: salt,
	}
}

type UserUsecase struct {
	store database.UserStoreInterface
	salt string
}

func (uu UserUsecase) CreateUser(ctx context.Context, user domain.UserData, password string) (*domain.User, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("Creating user with email: %", user.Email)
	hashedPassword := uu.HashPassword(password)
	user.Hash = hashedPassword
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

func (uu UserUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uu.store.GetUserByEmail(ctx, email)
}

func (uu UserUsecase) HashPassword(password string) string {
	passordBytes := []byte(password)
	hash := sha512.New()
	saltBytes := []byte(uu.salt)
	passordBytes = append(passordBytes, saltBytes...)
	hash.Write(passordBytes)
	hashedPasswordBytes := hash.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}

