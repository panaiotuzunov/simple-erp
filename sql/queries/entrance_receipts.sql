-- name: CreateEntranceReceipt :one
INSERT INTO entrance_receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, grain_type
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *, (gross - tare)::NUMERIC(12,3) AS net;

-- name: GetAllEntranceReceipts :many
SELECT *, (gross - tare)::NUMERIC(12,3) AS net
FROM entrance_receipts;

-- name: GetEntranceReceiptByID :one
SELECT *, (gross - tare)::NUMERIC(12,3) AS net 
FROM entrance_receipts
WHERE id = $1;



