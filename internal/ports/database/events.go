package database

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/FenceLive/public/model"
	"FenceLive/internal/ports/database/gen/FenceLive/public/table"
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
)

type EventStoreInterface interface {
	CreateEvent(ctx context.Context, event domain.EventData, tournamentId int64) (*domain.Event, error)
	GetByTournamentId(ctx context.Context, tournamentId int64) ([]*domain.Event, error)
}

type EventDatabaseStore struct {
	DB *sql.DB
}

func NewEventDatabaseStore(db *sql.DB) EventDatabaseStore {
	return EventDatabaseStore{
		DB: db,
	}
}

func (edbs EventDatabaseStore) CreateEvent(ctx context.Context, event domain.EventData, tournamentId int64) (*domain.Event, error) {
	modelEvent := model.Events{
		Name:         event.Name,
		Description:  event.Description,
		TournamentID: int32(tournamentId),
		StartTime:    event.Start,
		EndTime:      event.End,
		Status:       "CREATED",
		Weapon:       event.Weapon,
		Type:         event.Type,
		Category:     event.Category,
		Gender:       event.Gender,
	}

	stmt := table.Events.INSERT(table.Events.Name, table.Events.Description, table.Events.TournamentID, table.Events.StartTime, table.Events.EndTime, table.Events.Status, table.Events.Weapon, table.Events.Category, table.Events.Type, table.Events.Gender).
		MODEL(modelEvent).
		RETURNING(table.Events.AllColumns)

	var dest struct {
		model.Events
	}

	err := stmt.Query(edbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return &domain.Event{
		ID:           int64(dest.ID),
		Status:       dest.Status,
		TournamentId: int64(dest.TournamentID),
		EventData: domain.EventData{
			Name:        dest.Name,
			Type:        dest.Type,
			Gender:      dest.Gender,
			Category:    dest.Category,
			Description: dest.Description,
			Weapon:      dest.Weapon,
			Start:       dest.StartTime,
			End:         dest.EndTime,
		},
	}, nil
}

func (edbs EventDatabaseStore) GetByTournamentId(ctx context.Context, tournamentId int64) ([]*domain.Event, error) {
	stmt := table.Events.SELECT(table.Events.AllColumns).
		WHERE(table.Events.TournamentID.EQ(postgres.Int(tournamentId)))

	var dest []struct {
		model.Events
	}

	err := stmt.Query(edbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	var events []*domain.Event

	for _, evnt := range dest {
		events = append(events, &domain.Event{
			ID:           int64(evnt.ID),
			Status:       evnt.Status,
			TournamentId: int64(evnt.TournamentID),
			EventData: domain.EventData{
				Name:        evnt.Name,
				Description: evnt.Description,
				Start:       evnt.StartTime,
				End:         evnt.EndTime,
				Weapon:      evnt.Weapon,
				Type:        evnt.Type,
				Gender:      evnt.Gender,
				Category:    evnt.Category,
			},
		})
	}
	return events, nil
}
