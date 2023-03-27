package auth

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/usecases"
	"FenceLive/internal/usecases/users"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type AuthUsecase struct {
	JWTSecret     string
	JWTExpiration time.Duration
	Users         usecases.UserUsecaseInterface
}

func NewAuthUsecase(uu users.UserUsecase, jwtSecret string, JWTExpiration time.Duration) AuthUsecase {
	return AuthUsecase{
		JWTSecret:     jwtSecret,
		JWTExpiration: JWTExpiration,
		Users:         uu,
	}
}

func (au AuthUsecase) Login(ctx context.Context, creds domain.LoginCreds) (string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	usr, err := au.Users.GetUserByEmail(ctx, creds.Email)
	sugar.Infof("Attempting to log in with email: %s, and password: %s", creds.Email, creds.Password)
	if err != nil {
		sugar.Infof("Error whilst fetching from database")
		return "failure", err
	}
	hashedPassword := au.Users.HashPassword(creds.Password)
	if hashedPassword != usr.Hash {
		sugar.Infof("Password Missmatch")
		return "You suck", domain.InvalidCredentials
	}
	sugar.Infof("Credentials match")
	token, err := au.CreateJWT(ctx, *usr)
	if err != nil {
		return "skill issue", err
	}
	return token, nil
}

func (au AuthUsecase) CreateJWT(ctx context.Context, usr domain.User) (string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("Creating JWT token")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = usr.ID
	claims["username"] = usr.Username
	//TODO: add this
	// claims["role"] = usr.Role 
	claims["exp"] = time.Now().UTC().Add(au.JWTExpiration).Unix()

	resultToken, err := token.SignedString([]byte(au.JWTSecret))
	if err != nil {
		sugar.Errorf("Error creating JWT token")
		return "", err
	}
	sugar.Infof("Successfully created JWT token")
	return resultToken, nil
}
