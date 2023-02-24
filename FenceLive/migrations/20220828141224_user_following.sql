-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_following (
    follower_id INTEGER NOT NULL,
    following_id INTEGER NOT NULL,
    CONSTRAINT "user_following_pkey" PRIMARY KEY ("follower_id", "following_id"),
    CONSTRAINT "user_id_fk" FOREIGN KEY ("follower_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "following_id_fk" FOREIGN KEY ("following_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_following CASCADE;
-- +goose StatementEnd
