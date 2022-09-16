package database

import (
	"FenceLive/internal/domain"
	_ "FenceLive/migrations"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "FenceLive/migrations"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func createDockerUserStore(t *testing.T, ctx context.Context) (*UserDatabaseStore, func(), error) {
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
		return nil, nil, err
	}
	endpoint, err := postgresC.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	databaseUrl := fmt.Sprintf("postgres://test:test@%s/FenceLive?sslmode=disable", endpoint)
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Failed to open Db: %s", err)
		return nil, nil, err
	}
	err = goose.Up(db, "../../../migrations")
	if err != nil {
		log.Fatalf("Could not run migration: %s", err)
		return nil, nil, err
	}
	store := NewUserDatabaseStore(db)
	return store, func() {
		db.Close()
		postgresC.Terminate(ctx)
	}, nil
}

func TestUserDatabaseStore_Create(t *testing.T) {
	date := "2005-11-16"
	timeDate, _ := time.Parse("2006-01-02", date)
	type args struct {
		data domain.UserData
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		WantErr bool
	}{
		{
			name: "Create user",
			args: args{
				domain.UserData{
					BornIn:      timeDate,
					Email:       "maxelpicus@admin.com",
					Username:    "Maxek",
					FirstName:   "Max",
					LastName:    "Elpicus",
					Hash:        "8728u3ijldasd983oasd1pas9d",
					Nationality: "CZK",
				},
			},
			want: domain.User{
				ID: 1,
				UserData: domain.UserData{
					BornIn:      timeDate,
					Email:       "maxelpicus@admin.com",
					Username:    "Maxek",
					FirstName:   "Max",
					LastName:    "Elpicus",
					Hash:        "8728u3ijldasd983oasd1pas9d",
					Nationality: "CZK",
				},
			},
			WantErr: false,
		},
		{
			name: "create another user",
			args: args{
				domain.UserData{
					BornIn:      timeDate,
					Email:       "dapadxdxd@gmail.com",
					Username:    "espeletia",
					FirstName:   "Diego",
					LastName:    "Portillo",
					Hash:        "dsao8d09u2ejsd81o0sj0ask$2s",
					Nationality: "VEN",
				},
			},
			want: domain.User{
				ID: 2,
				UserData: domain.UserData{
					BornIn:      timeDate,
					Email:       "dapadxdxd@gmail.com",
					Username:    "espeletia",
					FirstName:   "Diego",
					LastName:    "Portillo",
					Hash:        "dsao8d09u2ejsd81o0sj0ask$2s",
					Nationality: "VEN",
				},
			},
			WantErr: false,
		},
		{
			name: "create user with same email",
			args: args{
				domain.UserData{
					BornIn:      timeDate,
					Email:       "dapadxdxd@gmail.com",
					Username:    "Diegx-PA",
					FirstName:   "Diego",
					LastName:    "Portillo",
					Hash:        "dsao8d09u2ejsd81o0sj0ask$2s",
					Nationality: "VEN",
				},
			},
			want:    domain.User{},
			WantErr: true,
		},
	}
	ctx := context.Background()
	users, cleanup, err := createDockerUserStore(t, ctx)
	if err != nil {
		t.Errorf("SpotDatabaseStore.Create() error = %v", err)
		return
	}
	t.Cleanup(cleanup)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := users.CreateUser(ctx, &tt.args.data)
			if err != nil {
				if !tt.WantErr {
					t.Errorf("Got an unwanted error: %s", err)
				}
			}
			if !tt.WantErr {
				AssertEqualUser(t, tt.want, *got)
			}
		})
	}
}

func AssertEqualUser(t testing.TB, want, got domain.User) {
	t.Helper()
	if want.BornIn.Unix() != got.BornIn.Unix() {
		t.Errorf("Got: %s, Wanted: %s", got.BornIn, want.BornIn)
	}
	if want.ID != got.ID {
		t.Errorf("Got: %d, Wanted: %d", got.ID, want.ID)
	}
	if want.Email != got.Email {
		t.Errorf("Got: %s, Wanted: %s", got.Email, want.Email)
	}
	if want.FirstName != got.FirstName {
		t.Errorf("Got: %s, Wanted: %s", got.FirstName, want.Email)
	}
	if want.LastName != got.LastName {
		t.Errorf("Got: %s, Wanted: %s", got.LastName, want.LastName)
	}
	if want.Username != got.Username {
		t.Errorf("Got: %s, Wanted: %s", got.Username, want.Username)
	}
	if want.Nationality != got.Nationality {
		t.Errorf("Got: %s, Wanted: %s", got.Nationality, want.Nationality)
	}
	if want.Hash != got.Hash {
		t.Errorf("Got: %s, Wanted: %s", got.Hash, want.Hash)
	}
}
