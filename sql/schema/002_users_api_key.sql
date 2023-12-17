-- +goose Up

ALTER TABLE users ADD COLUMN api_key varchar(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;