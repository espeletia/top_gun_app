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

func NewEventDatabaseStore(db *sql.DB) *EventDatabaseStore {
	return &EventDatabaseStore{
		DB: db,
	}
}

type EventDatabaseStore struct {
	DB *sql.DB
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
	var athletes []*domain.EventUser
	//Done: the shit has been fixed :)
	for _, athlete := range event.Athletes {
		athleteStmt := table.UserEvent.INSERT(table.UserEvent.UserID, table.UserEvent.EventID, table.UserEvent.PooleSeeding, table.UserEvent.Status).
			VALUES(athlete.UserID, dest.Events.ID, athlete.PooleSeeding, domain.AthleteCompeting).
			RETURNING(table.UserEvent.AllColumns)
		var athleteDest struct {
			model.UserEvent
		}
		err := athleteStmt.Query(edbs.DB, &athleteDest)
		if err != nil {
			return nil, err
		}
		athletes = append(athletes, mapDBEventUser(athleteDest.UserEvent))
	}

	storedEvent := mapDBEvent(dest.Events)
	storedEvent.Athletes = athletes

	return storedEvent, nil
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
		events = append(events, mapDBEvent(evnt.Events))
	}
	return events, nil
}

func mapDBEvent(Event model.Events) *domain.Event {
	return &domain.Event{
		ID:           int64(Event.ID),
		Status:       Event.Status,
		TournamentId: int64(Event.TournamentID),
		EventData: domain.EventData{
			Name:        Event.Name,
			Description: Event.Description,
			Start:       Event.StartTime,
			End:         Event.EndTime,
			Weapon:      Event.Weapon,
			Type:        Event.Type,
			Gender:      Event.Gender,
			Category:    Event.Category,
		},
	}
}

func mapDBEventUser(user model.UserEvent) *domain.EventUser {
	pooleSeeding := int64(*user.PooleSeeding)
	tableauSeeding := int64(*user.TableauSeeding)
	FinalRanking := int64(*user.FinalRanking)
	return &domain.EventUser{
		UserID:         int64(user.UserID),
		Status:         user.Status,
		Role:           user.UserRole,
		PooleSeeding:   &pooleSeeding,
		TableauSeeding: &tableauSeeding,
		FinalRanking:   &FinalRanking,
	}
}
