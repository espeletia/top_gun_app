-- +goose Up
-- +goose StatementBegin
CREATE TABLE tableau(
    id SERIAL NOT NULL,
    event_id SERIAL NOT NULL,
    status varchar(255) NOT NULL,
    
    CONSTRAINT "poole_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "tournament_id_fk" FOREIGN KEY ("tournament_id") REFERENCES "tournaments"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
