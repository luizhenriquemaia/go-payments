-- +goose Up
ALTER TABLE expense DROP COLUMN method;
ALTER TABLE expense DROP COLUMN account;
ALTER TABLE expense DROP COLUMN receipt;
ALTER TABLE expense DROP COLUMN paid_at;

-- +goose Down
ALTER TABLE expense
ADD COLUMN method SMALLINT NOT NULL DEFAULT -1;
ALTER TABLE expense
ADD COLUMN account SMALLINT NOT NULL DEFAULT -1;
ALTER TABLE expense
ADD COLUMN receipt VARCHAR(50) NOT NULL DEFAULT '';
ALTER TABLE expense
ADD COLUMN paid_at DATETIME NULL;
