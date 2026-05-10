-- name: GetStore :one
SELECT id, name FROM stores WHERE id = $1;

-- name: GetStoreByName :one
SELECT id, name FROM stores WHERE name = $1;

-- name: ListStores :many
SELECT id, name FROM stores ORDER BY name ASC;

-- name: ListStoresAsc :many
SELECT id, name FROM stores ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: ListStoresDesc :many
SELECT id, name FROM stores ORDER BY name DESC LIMIT $1 OFFSET $2;

-- name: ListStoresCount :one
SELECT COUNT(*) as count FROM stores;

-- name: CreateStore :one
INSERT INTO stores (name) VALUES ($1) RETURNING id, name;

-- name: UpdateStore :one
UPDATE stores SET name = $1 WHERE id = $2 RETURNING id, name;

-- name: DeleteStore :exec
DELETE FROM stores WHERE id = $1;
