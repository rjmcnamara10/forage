-- name: GetInventoryItem :one
SELECT id, item_id, unit_id, stored_amount FROM inventory_items WHERE id = $1;

-- name: GetInventoryItemByItemId :one
SELECT id, item_id, unit_id, stored_amount FROM inventory_items WHERE item_id = $1;

-- name: ListInventoryItems :many
SELECT id, item_id, unit_id, stored_amount FROM inventory_items ORDER BY item_id ASC LIMIT $1 OFFSET $2;

-- name: ListInventoryItemsCount :one
SELECT COUNT(*) as count FROM inventory_items;

-- name: CreateInventoryItem :one
INSERT INTO inventory_items (item_id, unit_id, stored_amount) VALUES ($1, $2, $3) RETURNING id, item_id, unit_id, stored_amount;

-- name: UpdateInventoryItem :one
UPDATE inventory_items SET unit_id = $1, stored_amount = $2 WHERE item_id = $3 RETURNING id, item_id, unit_id, stored_amount;

-- name: UpdateInventoryStoredAmount :one
UPDATE inventory_items SET stored_amount = $1 WHERE item_id = $2 RETURNING id, item_id, unit_id, stored_amount;

-- name: DeleteInventoryItem :exec
DELETE FROM inventory_items WHERE item_id = $1;

-- name: GetInventoryItemsWithItemInfo :many
SELECT ii.id, ii.item_id, i.name, ii.unit_id, u.name as unit_name, ii.stored_amount 
FROM inventory_items ii
JOIN items i ON ii.item_id = i.id
LEFT JOIN units u ON ii.unit_id = u.id
ORDER BY i.name ASC;

-- name: GetInventoryItemWithItemInfo :one
SELECT ii.id, ii.item_id, i.name, ii.unit_id, u.name as unit_name, ii.stored_amount 
FROM inventory_items ii
JOIN items i ON ii.item_id = i.id
LEFT JOIN units u ON ii.unit_id = u.id
WHERE ii.item_id = $1;
