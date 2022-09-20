-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
  id SERIAL NOT NULL,
  name varchar(255) NOT NULL,
  description TEXT,
  tournament_id SERIAL NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  status varchar(255) NOT NULL,
  weapon varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  gender varchar(255) NOT NULL,
  category varchar(255) NOT NULL,

  CONSTRAINT "event_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "tournament_fk" FOREIGN KEY ("tournament_id") REFERENCES "tournaments"("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events CASCADE;
-- +goose StatementEnd
