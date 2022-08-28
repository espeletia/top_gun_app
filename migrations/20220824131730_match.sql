-- +goose Up
-- +goose StatementBegin
CREATE TABLE match (
    id SERIAL NOT NULL,
    left_athlete_id SERIAL NOT NULL,
    right_athlete_id SERIAL NOT NULL,
    referee_id SERIAL NOT NULL,
    right_score INTEGER NOT NULL,
    left_score INTEGER NOT NULL,
    status varchar(255) NOT NULL,

    CONSTRAINT "poole_match_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "left_athlete_id_fk" FOREIGN KEY ("left_athlete_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "right_athlete_id_fk" FOREIGN KEY ("right_athlete_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "referee_id_fk" FOREIGN KEY ("referee_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "poole_id_fk" FOREIGN KEY ("poole_id") REFERENCES "poole"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
