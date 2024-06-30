-- +goose Up
ALTER TABLE expense
DROP COLUMN method,
DROP COLUMN account,
DROP COLUMN receipt,
DROP COLUMN paid_at;

-- +goose Down
ALTER TABLE expense
ADD COLUMN method SMALLINT NOT NULL DEFAULT(-1),
ADD COLUMN account SMALLINT NOT NULL DEFAULT(-1),
ADD COLUMN receipt VARCHAR(50) NOT NULL DEFAULT(''),
ADD COLUMN paid_at TIMESTAMP NULL;