-- name: CreateExitReceipt :one
INSERT INTO exit_receipts (
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

-- name: GetAllExitReceipts :many
SELECT * FROM exit_receipts;

-- name: GetExitReceiptByID :one
SELECT * FROM exit_receipts
WHERE id = $1;



