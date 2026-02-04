-- name: CreateEntranceReceipt :one
INSERT INTO entrance_receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, grain_type, purchase_id, company_id
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

-- name: GetAllEntranceReceipts :many
SELECT e.*,
    (e.gross - e.tare)::NUMERIC(12,3) AS net,
    p.suplier AS suplier
FROM entrance_receipts e
INNER JOIN purchases p
ON e.purchase_id = p.id;

-- name: GetEntranceReceiptByID :one
SELECT *, (gross - tare)::NUMERIC(12,3) AS net 
FROM entrance_receipts
WHERE id = $1;



