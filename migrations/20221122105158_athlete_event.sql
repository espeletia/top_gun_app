-- +goose Up
-- +goose StatementBegin
CREATE TABLE athlete_event (
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    initial_seeding integer NOT NULL,
    athlete_name varchar(20),
    athlete_last_name varchar(20),
    tableau_seeding integer,
    final_ranking integer,
    status varchar(255) NOT NULL,

    CONSTRAINT "athlete_tournament_pkey" PRIMARY KEY ("user_id", "event_id"),
    CONSTRAINT "user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events"("id") ON DELETE RESTRICT ON UPDATE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS athlete_event CASCADE;
-- +goose StatementEnd
