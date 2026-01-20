-- +goose Up
CREATE TABLE exit_receipts (
    id INT PRIMARY KEY DEFAULT nextval('receipt_id_seq'),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    truck_reg TEXT NOT NULL,
    trailer_reg TEXT NOT NULL,
    gross NUMERIC(12, 3) NOT NULL,
    tare NUMERIC(12, 3) NOT NULL,
    grain_type TEXT NOT NULL
);

-- +goose Down
DROP TABLE exit_receipts;