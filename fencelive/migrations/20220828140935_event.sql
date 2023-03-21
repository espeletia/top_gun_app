-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
  id SERIAL NOT NULL,
  name varchar(32) NOT NULL,
  description varchar(1024),
  tournament_id INTEGER NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  status varchar(16) NOT NULL,
  weapon varchar(16) NOT NULL,
  type varchar(16) NOT NULL,
  gender varchar(16) NOT NULL,
  category varchar(16) NOT NULL,

  CONSTRAINT "event_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "tournament_fk" FOREIGN KEY ("tournament_id") REFERENCES "tournaments"("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events CASCADE;
-- +goose StatementEnd
