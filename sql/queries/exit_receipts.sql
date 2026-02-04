-- name: CreateExitReceipt :one
INSERT INTO exit_receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, grain_type, sale_id, company_id
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *, (gross - tare)::NUMERIC(12,3) AS net;

-- name: GetAllExitReceipts :many
SELECT *, (gross - tare)::NUMERIC(12,3) AS net
FROM exit_receipts;

-- name: GetExitReceiptByID :one
SELECT *, (gross - tare)::NUMERIC(12,3) AS net 
FROM exit_receipts
WHERE id = $1;



