-- name: CreateTransport :one
INSERT INTO transports (
    created_at, updated_at, truck_reg, trailer_reg, net, sale_id, purchase_id, suplier_id, client_id
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
RETURNING *;

-- name: GetAllTransports :many
SELECT t.*, sc.name AS suplier, cc.name AS client
FROM transports t
INNER JOIN purchases p ON t.purchase_id = p.id
INNER JOIN sales s ON t.sale_id = s.id
INNER JOIN companies sc ON p.suplier_id = sc.id
INNER JOIN companies cc ON s.client_id = cc.id;