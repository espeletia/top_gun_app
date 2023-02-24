-- +goose Up
-- +goose StatementBegin
CREATE TABLE tableau(
    id SERIAL NOT NULL,
    name varchar(255) NOT NULL,
    event_id SERIAL NOT NULL,
    status varchar(255) NOT NULL,
    
    CONSTRAINT "tableau_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "tournament_id_fk" FOREIGN KEY ("event_id") REFERENCES "events"("id") ON DELETE RESTRICT ON UPDATE CASCADE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tableau CASCADE;
-- +goose StatementEnd
