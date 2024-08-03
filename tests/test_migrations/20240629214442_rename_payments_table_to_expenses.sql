-- +goose Up
ALTER TABLE payment
RENAME TO expense;

-- +goose Down
ALTER TABLE expense
RENAME TO payment;