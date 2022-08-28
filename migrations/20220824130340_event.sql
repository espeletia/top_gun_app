-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
  id SERIAL NOT NULL,
  name varchar(255) NOT NULL,
  description TEXT,
  tournament_id SERIAL NOT NULL,
  start timestamp NOT NULL,
  end timestamp NOT NULL,
  status varchar(255),
  weapon varchar(255),
  type varchar(255),
  gender varchar(255),
  category varchar(255),

  CONSTRAINT "event_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "tournament_fk" FOREIGN KEY ("tournament_id") REFERENCES "tournaments"("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
