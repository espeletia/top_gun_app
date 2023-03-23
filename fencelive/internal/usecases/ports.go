package usecases

type HashUsecaseInterface interface {
	HashPassword(password string) string
}
