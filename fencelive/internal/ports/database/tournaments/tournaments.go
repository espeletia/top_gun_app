package tournaments

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/postgres/public/model"
	"FenceLive/internal/ports/database/gen/postgres/public/table"
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
)

func NewTournamentDatabaseStore(db *sql.DB) *TournamentDatabaseStore {
	return &TournamentDatabaseStore{
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
		City:        TournData.City,
		Status:      domain.TournamentStatusCreated,
		OwnerID:     int32(TournData.OwnerId),
		Country:     TournData.Country,
		Description: TournData.Description,
	}

	if TournData.Location != nil {
		modelTourn.Lat = &TournData.Location.Lat
		modelTourn.Lon = &TournData.Location.Lon
		modelTourn.Address = &TournData.Location.Address
	}

	stmt := table.Tournaments.INSERT(table.Tournaments.StartTime, table.Tournaments.EndTime, table.Tournaments.Name, table.Tournaments.Lat, table.Tournaments.Lon, table.Tournaments.Address, table.Tournaments.City, table.Tournaments.Country, table.Tournaments.OwnerID, table.Events.Description, table.Tournaments.Status).
		MODEL(modelTourn).
		RETURNING(table.Tournaments.AllColumns)

	var dest struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)

	if err != nil {
		return nil, err
	}
	return mapDBTournament(dest.Tournaments), nil
}

func (tdbs TournamentDatabaseStore) UpdateTournamentData(ctx context.Context, tournamentID int64, TournData domain.TournamentData) (*domain.Tournament, error) {
	modelTourn := model.Tournaments{
		StartTime:   TournData.Start,
		EndTime:     TournData.End,
		Name:        TournData.Name,
		City:        TournData.City,
		OwnerID:     int32(TournData.OwnerId),
		Country:     TournData.Country,
		Description: TournData.Description,
	}

	stmt := table.Tournaments.UPDATE(table.Tournaments.StartTime, table.Tournaments.EndTime, table.Tournaments.Name, table.Tournaments.City, table.Tournaments.OwnerID, table.Tournaments.Country, table.Tournaments.Description).
		MODEL(modelTourn).WHERE(table.Tournaments.ID.EQ(postgres.Int(tournamentID))).RETURNING(table.Tournaments.AllColumns)

	var dest struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return mapDBTournament(dest.Tournaments), nil
}

func (tdbs TournamentDatabaseStore) GetTournamentById(ctx context.Context, id int64) (*domain.Tournament, error) {
	stmt := table.Tournaments.SELECT(table.Tournaments.AllColumns).WHERE(table.Tournaments.ID.EQ(postgres.Int(id)))

	var dest []struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	if len(dest) < 1 {
		return nil, domain.TournamentNotFound
	}

	return mapDBTournament(dest[0].Tournaments), nil
}

func (tdbs TournamentDatabaseStore) GetAllTournaments(ctx context.Context) ([]*domain.Tournament, error) {
	stmt := table.Tournaments.SELECT(table.Tournaments.AllColumns)

	var dest []struct {
		model.Tournaments
	}

	err := stmt.Query(tdbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	var tournaments []*domain.Tournament
	for _, tournament := range dest {
		tournaments = append(tournaments, mapDBTournament(tournament.Tournaments))
	}

	return tournaments, nil
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
			City:        tournament.City,
			OwnerId:     int64(tournament.OwnerID),
			Location:    loc,
		},
	}
}
