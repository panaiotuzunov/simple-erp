-- name: CreateCompany :one
INSERT INTO companies (
    created_at,
    updated_at,
    vat_number,
    name,
    address
)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetCompanyByID :one
SELECT * FROM companies
WHERE id = $1;

-- name: GetCompanyByVATNumber :one
SELECT * FROM companies
WHERE vat_number = $1;

-- name: GetCompanyByName :one
SELECT * FROM companies
WHERE name = $1;

-- name: GetAllCompanies :many
SELECT * FROM companies
ORDER BY id;