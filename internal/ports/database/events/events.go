package events

import (
	"FenceLive/internal/domain"
	"FenceLive/internal/ports/database/gen/FenceLive/public/model"
	"FenceLive/internal/ports/database/gen/FenceLive/public/table"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
)

func NewEventDatabaseStore(db *sql.DB) *EventDatabaseStore {
	return &EventDatabaseStore{
		DB: db,
	}
}

type EventDatabaseStore struct {
	DB *sql.DB
}

func (edbs EventDatabaseStore) GetEventById(ctx context.Context, eventId int64) (*domain.Event, error) {
	stmt := table.Events.SELECT(table.Events.AllColumns).WHERE(table.Events.ID.EQ(postgres.Int(eventId)))
	var dest []struct {
		model.Events
	}

	err := stmt.Query(edbs.DB, &dest)

	if err != nil {
		return nil, err
	}
	if len(dest) < 1 {
		return nil, domain.EventNotFound
	}

	return mapDBEvent(dest[0].Events), nil
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
	var athletes []*domain.Athlete
	//Done: the shit has been fixed :)
	//nvm
	fmt.Printf("athletes: %v\n", event.Athletes)
	for _, athlete := range event.Athletes {
		athleteStmt := table.AthleteEvent.INSERT(table.AthleteEvent.UserID, table.AthleteEvent.EventID, table.AthleteEvent.InitialSeeding, table.AthleteEvent.Status).
			VALUES(athlete.UserID, dest.Events.ID, athlete.PooleSeeding, domain.AthleteCompeting).
			RETURNING(table.AthleteEvent.AllColumns)
		var athleteDest struct {
			model.AthleteEvent
		}
		fmt.Printf("stored athletes: %v\n", athleteDest)
		err := athleteStmt.Query(edbs.DB, &athleteDest)
		if err != nil {
			return nil, err
		}
		athletes = append(athletes, mapDBEventAthlete(athleteDest.AthleteEvent))
		fmt.Printf("appended athletes: %v\n", athleteDest)
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

func (edbs EventDatabaseStore) GetAllAthletes(ctx context.Context, eventId int64) ([]*domain.Athlete, error) {
	stmt := table.AthleteEvent.SELECT(table.AthleteEvent.AllColumns).
		WHERE(table.AthleteEvent.EventID.EQ(postgres.Int(eventId)))

	var dest []struct {
		model.AthleteEvent
	}

	err := stmt.Query(edbs.DB, &dest)
	if err != nil {
		return nil, err
	}

	var athletes []*domain.Athlete
	for _, athlete := range dest {
		mappedAthlete := mapDBEventAthlete(athlete.AthleteEvent)
		athletes = append(athletes, mappedAthlete)
	}
	return athletes, nil
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

func mapDBEventAthlete(user model.AthleteEvent) *domain.Athlete {
	pooleSeeding := int64(user.InitialSeeding)
	var tableauSeeding *int64
	if user.TableauSeeding != nil {
		tableauSeedingval := int64(*user.TableauSeeding)
		tableauSeeding = &tableauSeedingval
	}
	var finalRanking *int64
	if user.FinalRanking != nil {
		finalRankingval := int64(*user.FinalRanking)
		finalRanking = &finalRankingval
	}
	return &domain.Athlete{
		UserID:         int64(user.UserID),
		Status:         user.Status,
		PooleSeeding:   pooleSeeding,
		TableauSeeding: tableauSeeding,
		FinalRanking:   finalRanking,
	}
}
