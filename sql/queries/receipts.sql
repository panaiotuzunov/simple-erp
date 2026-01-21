-- name: GetAllReceiptsUnion :many
SELECT 
    *,
    (gross - tare)::NUMERIC(12,3) AS net,
    'entrance' AS receipt_type 
FROM entrance_receipts

UNION ALL

SELECT 
    *,
    (gross - tare)::NUMERIC(12,3) AS net,
    'exit' AS receipt_type 
FROM exit_receipts
ORDER BY id;