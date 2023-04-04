package setup

import (
	"FenceLive/graph"
	"FenceLive/internal/config"
	"FenceLive/internal/ports/database/events"
	"FenceLive/internal/ports/database/tournaments"
	"FenceLive/internal/ports/database/users"
	"FenceLive/internal/usecases/auth"
	eventUsecase "FenceLive/internal/usecases/events"
	tournamentUsecase "FenceLive/internal/usecases/tournaments"
	userUsecase "FenceLive/internal/usecases/users"
	"database/sql"
)

func NewResolver(dbConn *sql.DB, config config.Config) (*graph.Resolver, error) {
	tournamentStore := tournaments.NewTournamentDatabaseStore(dbConn)
	tournamentUsecase := tournamentUsecase.NewTournamentUsecase(tournamentStore)
	eventStore := events.NewEventDatabaseStore(dbConn)
	eventUsecase := eventUsecase.NewEventUsecase(eventStore)
	userStore := users.NewUserDatabaseStore(dbConn)
	userUsecase := userUsecase.NewUserUsecase(userStore, config.HashConfig.Salt)
	authUsecase := auth.NewAuthUsecase(userUsecase, config.JWTConfig.Signature, config.JWTConfig.Expiration)

	return &graph.Resolver{
		Tournaments: tournamentUsecase,
		Events:      eventUsecase,
		Users:       userUsecase,
		Auth:        authUsecase,

		Mapper:      graph.NweGqlMapper(),
		InputMapper: graph.NewInputMapper(),
	}, nil

}
