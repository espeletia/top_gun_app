-- +goose Up
-- +goose StatementBegin
CREATE TABLE poole(
    id SERIAL NOT NULL,
    event_id SERIAL NOT NULL,
    referee_id SERIAL NOT NULL,
    status varchar(255) NOT NULL,
    
    CONSTRAINT "poole_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "tournament_id_fk" FOREIGN KEY ("tournament_id") REFERENCES "tournaments"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "referee_id_fk" FOREIGN KEY ("referee_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS poole;
-- +goose StatementEnd
