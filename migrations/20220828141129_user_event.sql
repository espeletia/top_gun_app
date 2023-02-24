-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_event (
    user_id SERIAL NOT NULL,
    event_id SERIAL NOT NULL,
    user_role VARCHAR(255) NOT NULL,
    poole_seeding integer,
    tableau_seeding integer,
    final_ranking integer,
    status varchar(255),

    CONSTRAINT "user_tournament_pkey" PRIMARY KEY ("user_id", "event_id"),
    CONSTRAINT "user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events"("id") ON DELETE RESTRICT ON UPDATE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_event CASCADE;
-- +goose StatementEnd
