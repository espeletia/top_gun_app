-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL NOT NULL,
  first_name varchar(32) NOT NULL,
  last_name varchar(32) NOT NULL,
  username varchar(16) UNIQUE NOT NULL,
  email varchar(50) UNIQUE NOT NULL,
  nationality varchar(32) NOT NULL,
  born_in date NOT NULL,
  hash varchar(256) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  role varchar(16),

  CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
