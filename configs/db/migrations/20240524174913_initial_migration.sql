-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment(
    id SERIAL PRIMARY KEY,
    description VARCHAR(100) NOT NULL,
    cost_center SMALLINT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE payment;

-- +goose StatementEnd