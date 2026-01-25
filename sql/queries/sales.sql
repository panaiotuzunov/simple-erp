-- name: CreateSale :one
INSERT INTO sales (
    created_at, updated_at, client, price, quantity, grain_type 
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
SELECT * FROM sales;