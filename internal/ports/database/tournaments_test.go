package database

import (
	"FenceLive/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func createDockerTournamentStore(t *testing.T, ctx context.Context) (*UserDatabaseStore, *TournamentDatabaseStore, func(), error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "FenceLive",
			"listen_addresses":  "'*'",
		},
		WaitingFor: wait.ForSQL(nat.Port("5432"), "postgres", func(p nat.Port) string {
			return fmt.Sprintf("postgres://test:test@localhost:%v/FenceLive?sslmode=disable", p.Port())
		}),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	endpoint, err := postgresC.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, err
	}
	databaseUrl := fmt.Sprintf("postgres://test:test@%s/FenceLive?sslmode=disable", endpoint)
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Failed to open Db: %s", err)
		return nil, nil, nil, err
	}
	err = goose.Up(db, "../../../migrations")
	if err != nil {
		log.Fatalf("Could not run migration: %s", err)
		return nil, nil, nil, err
	}
	store := NewUserDatabaseStore(db)
	return store, NewTournamentDatabaseStore(db), func() {
		db.Close()
		postgresC.Terminate(ctx)
	}, nil
}

func TestTournamentDatabaseStore_Create(t *testing.T) {
	dateFormat := "2005-11-16"
	dateStart, _ := time.Parse("2006-01-05", dateFormat)
	dateEnd, _ := time.Parse("2006-01-06", dateFormat)
	type args struct {
		data domain.TournamentData
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Tournament
		WantErr bool
	}{
		{
			name: "Create Tournament",
			args: args{
				domain.TournamentData{
					Start:       dateStart,
					End:         dateEnd,
					City:        "Hradec Kralove",
					Name:        "Karluv memorial",
					Country:     "CZE",
					OwnerId:     1,
					Description: strptr("memorial WWII"),
				},
			},
			want: domain.Tournament{
				Id:     1,
				Status: domain.TournamentStatusCreated,
				TournamentData: domain.TournamentData{
					Start:       dateStart,
					End:         dateEnd,
					City:        "Hradec Kralove",
					Name:        "Karluv memorial",
					Country:     "CZE",
					OwnerId:     1,
					Description: strptr("memorial WWII"),
					Location:    nil,
				},
			},
			WantErr: false,
		},
	}
	ctx := context.Background()
	users, tournaments, cleanup, err := createDockerTournamentStore(t, ctx)
	if err != nil {
		t.Errorf("UserDatabaseStore.Create() error = %v", err)
		return
	}
	t.Cleanup(cleanup)
	_, _ = users.CreateUser(ctx, domain.UserData{
		BornIn:      dateStart,
		Email:       "maxelpicus@admin.com",
		Username:    "Maxek",
		FirstName:   "Max",
		LastName:    "Elpicus",
		Hash:        "8728u3ijldasd983oasd1pas9d",
		Nationality: "CZK",
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tournaments.CreateTournament(ctx, tt.args.data)
			if err != nil {
				if !tt.WantErr {
					t.Errorf("Got an unwanted error: %s", err)
				}
			}
			if !tt.WantErr {
				assertEqualTournament(t, tt.want, *got)
			}
		})
	}

}

func assertEqualTournament(t testing.TB, want, got domain.Tournament) {
	t.Helper()
	if want.Start.Unix() != got.Start.Unix() {
		t.Errorf("Got: %s, Wanted: %s", got.Start, want.Start)
	}
	if want.End.Unix() != got.End.Unix() {
		t.Errorf("Got: %s, Wanted: %s", got.End, want.End)
	}
	if want.Id != got.Id {
		t.Errorf("Got: %d, Wanted: %d", got.Id, want.Id)
	}
	if want.OwnerId != got.OwnerId {
		t.Errorf("Got: %d, Wanted: %d", got.OwnerId, want.OwnerId)
	}
	if *want.Description != *got.Description {
		t.Errorf("Got: %s, Wanted: %s", *got.Description, *want.Description)
	}
	if want.Name != got.Name {
		t.Errorf("Got: %s, Wanted: %s", got.Name, want.Name)
	}
	if want.Status != got.Status {
		t.Errorf("Got: %s, Wanted: %s", got.Status, want.Status)
	}
	if want.OwnerId != got.OwnerId {
		t.Errorf("Got: %d, Wanted: %d", got.OwnerId, want.OwnerId)
	}
	if want.Country != got.Country {
		t.Errorf("Got: %s, Wanted: %s", got.Country, want.Country)
	}
	if want.City != got.City {
		t.Errorf("Got: %s, Wanted: %s", got.City, want.City)
	}
	if want.Location != nil {
		if got.Location != nil {
			if want.Location.Lat != got.Location.Lat {
				t.Errorf("Got: %f, Wanted: %f", got.Location.Lat, want.Location.Lat)
			}
			if want.Location.Lon != got.Location.Lon {
				t.Errorf("Got: %f, Wanted: %f", got.Location.Lon, want.Location.Lon)
			}
			if want.Location.Address != got.Location.Address {
				t.Errorf("Got: %s, Wanted: %s", got.Location.Address, want.Location.Address)
			}
		}
	}
}

func strptr(s string) *string {
	return &s
}
