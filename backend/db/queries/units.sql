-- name: GetUnit :one
SELECT id, name FROM units WHERE id = $1;

-- name: GetUnitByName :one
SELECT id, name FROM units WHERE name = $1;

-- name: ListUnits :many
SELECT id, name FROM units ORDER BY name ASC;

-- name: CreateUnit :one
INSERT INTO units (name) VALUES ($1) RETURNING id, name;

-- name: UpdateUnit :exec
UPDATE units SET name = $1 WHERE id = $2;

-- name: DeleteUnit :exec
DELETE FROM units WHERE id = $1;
