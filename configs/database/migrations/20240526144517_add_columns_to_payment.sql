-- +goose Up
ALTER TABLE payment
ADD COLUMN status SMALLINT NOT NULL DEFAULT(0),
ADD COLUMN bar_code VARCHAR(60) NOT NULL,
ADD COLUMN updated_at TIMESTAMP NOT NULL,
ADD COLUMN created_at TIMESTAMP NOT NULL;

-- +goose Down
ALTER TABLE payment 
DROP COLUMN status,
DROP COLUMN bar_code,
DROP COLUMN updated_at,
DROP COLUMN created_at;