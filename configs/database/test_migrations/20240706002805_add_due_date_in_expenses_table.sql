-- +goose Up
ALTER TABLE expense
ADD COLUMN due_date DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL;

-- +goose Down
ALTER TABLE expense
DROP COLUMN due_date;