-- +goose Up
-- +goose StatementBegin
CREATE TABLE tournaments (
  id SERIAL NOT NULL,
  owner_id SERIAL NOT NULL,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL,
  city varchar(255),
  country varchar(255),
  status varchar(255),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  lat float(53),
  lon float(53),
  address TEXT,

  CONSTRAINT "tournament_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tournaments CASCADE;
-- +goose StatementEnd