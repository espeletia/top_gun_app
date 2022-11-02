package database

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/FenceLive/public/model"
	"FenceLive/internal/ports/database/gen/FenceLive/public/table"
	"context"
	"database/sql"
	"fmt"
)

type TournamentStoreInterface interface {
	CreateTournament(ctx context.Context, TournData domain.TournamentData) (*domain.Tournament, error)
}

func NewTournamentDatabaseStore(db *sql.DB) *TournamentDatabaseStore {
	return &TournamentDatabaseStore{
		DB: db,
	}
}

type TournamentDatabaseStore struct {
	DB *sql.DB
}

func (tdbs TournamentDatabaseStore) CreateTournament(ctx context.Context, TournData domain.TournamentData) (*domain.Tournament, error) {
	fmt.Printf("??????????????//\n")
	modelTourn := model.Tournaments{
		StartTime:   TournData.Start,
		EndTime:     TournData.End,
		Name:        TournData.Name,
		City:        TournData.City,
		Status:      "CREATED",
		OwnerID:     int32(TournData.OwnerId),
		Country:     TournData.Country,
		Description: TournData.Description,
	}

	if TournData.Location != nil {
		modelTourn.Lat = &TournData.Location.Lat
		modelTourn.Lon = &TournData.Location.Lon
		modelTourn.Address = &TournData.Location.Address
	}

	fmt.Printf("we cool?\n")

	stmt := table.Tournaments.INSERT(table.Tournaments.StartTime, table.Tournaments.EndTime, table.Tournaments.Name, table.Tournaments.Lat, table.Tournaments.Lon, table.Tournaments.Address, table.Tournaments.City, table.Tournaments.Country, table.Tournaments.OwnerID, table.Events.Description, table.Tournaments.Status).
		MODEL(modelTourn).
		RETURNING(table.Tournaments.AllColumns)
	fmt.Printf("we cool\n")

	var dest struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)
	fmt.Printf("we cool %v\n", dest)

	if err != nil {
		return nil, err
	}

	// stored := &domain.Tournament{
	// 	Id:     int64(dest.ID),
	// 	Status: dest.Status,
	// 	TournamentData: domain.TournamentData{
	// 		Start:       dest.StartTime,
	// 		End:         dest.EndTime,
	// 		Name:        dest.Name,
	// 		City:        dest.City,
	// 		Country:     dest.Country,
	// 		OwnerId:     int64(dest.OwnerID),
	// 		Description: dest.Description,
	// 	},
	// }
	// fmt.Printf("we cool %f, %s, %f\n", *dest.Lat, *dest.Address, *dest.Lon)
	// fmt.Printf("we cool %p, %p, %p\n", dest.Lat, dest.Address, dest.Lon)

	// if dest.Lat != nil {
	// 	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n")
	// 	stored.Location = &domain.Location{
	// 		Lon:     *dest.Lon,
	// 		Lat:     *dest.Lat,
	// 		Address: *dest.Address,
	// 	}
	// }
	// fmt.Printf("we cooled\n")
	return mapDBTournament(dest.Tournaments), nil
}

func mapDBTournament(tournament model.Tournaments) *domain.Tournament {
	var loc *domain.Location
	if tournament.Lat != nil {
		loc = &domain.Location{
			Lon:     *tournament.Lon,
			Lat:     *tournament.Lat,
			Address: *tournament.Address,
		}
	}
	return &domain.Tournament{
		Id:     int64(tournament.ID),
		Status: tournament.Status,
		TournamentData: domain.TournamentData{
			Start:       tournament.StartTime,
			End:         tournament.EndTime,
			Name:        tournament.Name,
			Description: tournament.Description,
			Country:     tournament.Country,
			OwnerId:     int64(tournament.OwnerID),
			Location:    loc,
		},
	}
}
