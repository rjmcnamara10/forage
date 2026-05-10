-- name: GetItem :one
SELECT id, name FROM items WHERE id = $1;

-- name: GetItemByName :one
SELECT id, name FROM items WHERE name = $1;

-- name: ListItems :many
SELECT id, name FROM items ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: ListItemsCount :one
SELECT COUNT(*) as count FROM items;

-- name: CreateItem :one
INSERT INTO items (name) VALUES ($1) RETURNING id, name;

-- name: UpdateItem :one
UPDATE items SET name = $1 WHERE id = $2 RETURNING id, name;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1;

-- name: GetItemCategories :many
SELECT ic.id, ic.name FROM item_categories ic
JOIN item_category_mappings icm ON ic.id = icm.category_id
WHERE icm.item_id = $1
ORDER BY ic.name ASC;

-- name: AddItemCategory :exec
INSERT INTO item_category_mappings (item_id, category_id) VALUES ($1, $2);

-- name: RemoveItemCategory :exec
DELETE FROM item_category_mappings WHERE item_id = $1 AND category_id = $2;

-- name: RemoveAllItemCategories :exec
DELETE FROM item_category_mappings WHERE item_id = $1;

-- name: ListItemsByCategory :many
SELECT DISTINCT i.id, i.name FROM items i
JOIN item_category_mappings icm ON i.id = icm.item_id
WHERE icm.category_id = $1
ORDER BY i.name ASC LIMIT $2 OFFSET $3;

-- name: ListItemsByCategoryCount :one
SELECT COUNT(DISTINCT i.id) as count FROM items i
JOIN item_category_mappings icm ON i.id = icm.item_id
WHERE icm.category_id = $1;

-- name: SearchItems :many
SELECT id, name FROM items WHERE name ILIKE $1 ORDER BY name ASC LIMIT $2 OFFSET $3;

-- name: SearchItemsDesc :many
SELECT id, name FROM items WHERE name ILIKE $1 ORDER BY name DESC LIMIT $2 OFFSET $3;

-- name: SearchItemsCount :one
SELECT COUNT(*) as count FROM items WHERE name ILIKE $1;

-- name: ListItemsDesc :many
SELECT id, name FROM items ORDER BY name DESC LIMIT $1 OFFSET $2;

-- name: ListItemsByCategoryDesc :many
SELECT DISTINCT i.id, i.name FROM items i
JOIN item_category_mappings icm ON i.id = icm.item_id
WHERE icm.category_id = $1
ORDER BY i.name DESC LIMIT $2 OFFSET $3;
