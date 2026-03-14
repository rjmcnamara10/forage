-- name: GetStore :one
SELECT id, name FROM stores WHERE id = $1;

-- name: GetStoreByName :one
SELECT id, name FROM stores WHERE name = $1;

-- name: ListStores :many
SELECT id, name FROM stores ORDER BY name ASC;

-- name: CreateStore :one
INSERT INTO stores (name) VALUES ($1) RETURNING id, name;

-- name: UpdateStore :exec
UPDATE stores SET name = $1 WHERE id = $2;

-- name: DeleteStore :exec
DELETE FROM stores WHERE id = $1;
