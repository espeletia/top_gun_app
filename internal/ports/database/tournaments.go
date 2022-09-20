package database

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/FenceLive/public/model"
	"FenceLive/internal/ports/database/gen/FenceLive/public/table"
	"context"
	"database/sql"
)

type TournamentStoreInterface interface {
	CreateTournament(ctx context.Context, TournData domain.TournamentData) (*domain.Tournament, error)
}

func NewTournamentDatabaseStore(db *sql.DB) TournamentDatabaseStore {
	return TournamentDatabaseStore{
		DB: db,
	}
}

type TournamentDatabaseStore struct {
	DB *sql.DB
}

func (tdbs TournamentDatabaseStore) CreateTournament(ctx context.Context, TournData domain.TournamentData) (*domain.Tournament, error) {
	modelTourn := model.Tournaments{
		StartTime:   TournData.Start,
		EndTime:     TournData.End,
		Name:        TournData.Name,
		Lat:         &TournData.Location.Lat,
		Lon:         &TournData.Location.Lon,
		Address:     &TournData.Location.Address,
		City:        TournData.City,
		Status:      "CREATED",
		OwnerID:     int32(TournData.OwnerId),
		Country:     TournData.Country,
		Description: TournData.Description,
	}

	stmt := table.Tournaments.INSERT(table.Tournaments.StartTime, table.Tournaments.EndTime, table.Tournaments.Name, table.Tournaments.Lat, table.Tournaments.Lon, table.Tournaments.Address, table.Tournaments.City, table.Tournaments.Country, table.Tournaments.OwnerID, table.Events.Description).
		MODEL(modelTourn).
		RETURNING(table.Tournaments.AllColumns)

	var dest struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)

	if err != nil {
		return nil, err
	}
	var loc *domain.Location
	if dest.Lat != nil {
		loc.Address = *dest.Address
		loc.Lat = *dest.Lat
		loc.Lon = *dest.Lon
	}

	return &domain.Tournament{
		Id:     int64(dest.ID),
		Status: dest.Status,
		TournamentData: domain.TournamentData{
			Start:       dest.StartTime,
			End:         dest.EndTime,
			Name:        dest.Name,
			Location:    loc,
			City:        dest.City,
			Country:     dest.Country,
			OwnerId:     int64(dest.OwnerID),
			Description: dest.Description,
		},
	}, nil
}
