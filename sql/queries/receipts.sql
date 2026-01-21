-- name: GetAllReceiptsUnion :many
SELECT 
    *,
    (gross - tare)::NUMERIC(12,3) AS net,
    'entrance' AS receipt_type 
FROM entrance_receipts

UNION ALL

SELECT 
    id,
    created_at,
    updated_at,
    truck_reg,
    trailer_reg,
    (gross * -1)::NUMERIC(12,3) AS gross,
    (tare * -1)::NUMERIC(12,3) AS tare,
    grain_type,
    ((gross - tare) * -1)::NUMERIC(12,3) AS net,
    'exit' AS receipt_type 
FROM exit_receipts
ORDER BY id;