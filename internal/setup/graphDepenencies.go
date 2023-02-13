package setup

import (
	"FenceLive/graph"
	"FenceLive/internal/usecases"
)

func NewRsolver() (*graph.Resolver, error) {
	tournamentUsecase := usecases.NewTournamentUsecase()
}
