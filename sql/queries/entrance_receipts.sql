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
RETURNING *;

-- name: GetAllEntranceReceipts :many
SELECT * FROM entrance_receipts;

