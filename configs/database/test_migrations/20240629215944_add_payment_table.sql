-- +goose Up
CREATE TABLE IF NOT EXISTS payment(
    id SERIAL PRIMARY KEY,
    expense_id INT,
    receipt VARCHAR(50) NOT NULL DEFAULT(''),
    method SMALLINT NOT NULL DEFAULT(-1),
    account SMALLINT NOT NULL DEFAULT(-1),
    paid_at DATETIME NULL,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (expense_id) REFERENCES expense(id) ON DELETE NO ACTION
);

-- +goose Down
DROP TABLE payment;