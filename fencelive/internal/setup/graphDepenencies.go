package setup

import (
	"FenceLive/graph"
	"FenceLive/internal/config"
	"FenceLive/internal/ports/database/events"
	"FenceLive/internal/ports/database/tournaments"
	"FenceLive/internal/ports/database/users"
	eventUsecase "FenceLive/internal/usecases/events"
	"FenceLive/internal/usecases/hash"
	tournamentUsecase "FenceLive/internal/usecases/tournaments"
	userUsecase "FenceLive/internal/usecases/users"
	"database/sql"
)

func NewResolver(dbConn *sql.DB, config config.Config) (*graph.Resolver, error) {
	Hasher := hash.NewHashUsecase(config.HashConfig)
	tournamentStore := tournaments.NewTournamentDatabaseStore(dbConn)
	tournamentUsecase := tournamentUsecase.NewTournamentUsecase(tournamentStore)
	eventStore := events.NewEventDatabaseStore(dbConn)
	eventUsecase := eventUsecase.NewEventUsecase(eventStore)
	userStore := users.NewUserDatabaseStore(dbConn)
	userUsecase := userUsecase.NewUserUsecase(userStore, Hasher)

	return &graph.Resolver{
		Tournaments: tournamentUsecase,
		Events:      eventUsecase,
		Users:       userUsecase,

		Mapper:      graph.NweGqlMapper(),
		InputMapper: graph.NewInputMapper(),
	}, nil

}
