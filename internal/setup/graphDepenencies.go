package setup

import (
	"FenceLive/graph"
	"FenceLive/internal/ports/database/events"
	"FenceLive/internal/ports/database/tournaments"
	"FenceLive/internal/ports/database/users"
	"FenceLive/internal/usecases"
	"database/sql"
)

func NewResolver(dbConn *sql.DB) (*graph.Resolver, error) {
	tournamentStore := tournaments.NewTournamentDatabaseStore(dbConn)
	tournamentUsecase := usecases.NewTournamentUsecase(tournamentStore)
	eventStore := events.NewEventDatabaseStore(dbConn)
	eventUsecase := usecases.NewEventUsecase(eventStore)
	userStore := users.NewUserDatabaseStore(dbConn)
	userUsecase := usecases.NewUserUsecase(userStore)
	
	return &graph.Resolver{
		Tournaments: tournamentUsecase,
		Events: eventUsecase,
		Users: userUsecase,

		Mapper: graph.NweGqlMapper(),
		InputMapper: graph.NewInputMapper(),
	}, nil

}
