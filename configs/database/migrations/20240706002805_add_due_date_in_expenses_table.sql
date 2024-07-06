-- +goose Up
ALTER TABLE expense
ADD COLUMN due_date TIMESTAMP DEFAULT NOW() NOT NULL;

-- +goose Down
ALTER TABLE expense
DROP COLUMN due_date;
