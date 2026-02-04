-- name: CreateSale :one
INSERT INTO sales (
    created_at, updated_at, company_id, price, quantity, grain_type 
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetSaleById :one
SELECT * FROM sales
WHERE id = $1;

-- name: GetAllSales :many
SELECT s.*, 
    COALESCE(SUM(e.gross - e.tare), 0)::NUMERIC(12, 3) AS expedited_receipts,
    COALESCE(SUM(t.net), 0)::NUMERIC(12, 3) AS expedited_transports
FROM sales s
LEFT JOIN exit_receipts e
ON s.id = e.sale_id
LEFT JOIN transports t
ON s.id = t.sale_id
GROUP BY s.id
ORDER BY s.id;