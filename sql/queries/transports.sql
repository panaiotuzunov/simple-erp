-- name: CreateTransport :one
INSERT INTO transports (
    created_at, updated_at, truck_reg, trailer_reg, net, sale_id, purchase_id
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

-- name: GetAllTransports :many
SELECT t.*, p.suplier, s.client
FROM transports t
INNER JOIN purchases p
ON t.purchase_id = p.id
INNER JOIN sales s
ON t.sale_id = s.id;