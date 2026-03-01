-- name: GetStoreItem :one
SELECT id, store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order 
FROM store_items WHERE id = $1;

-- name: ListStoreItems :many
SELECT id, store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order 
FROM store_items ORDER BY store_id ASC, store_traversal_order ASC LIMIT $1 OFFSET $2;

-- name: ListStoreItemsCount :one
SELECT COUNT(*) as count FROM store_items;

-- name: CreateStoreItem :one
INSERT INTO store_items (store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING id, store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order;

-- name: UpdateStoreItem :exec
UPDATE store_items 
SET purchase_unit_id = $1, inventory_unit_id = $2, inventory_units_per_purchase = $3, purchased_by_decimal = $4, is_favorite = $5, store_traversal_order = $6 
WHERE id = $7;

-- name: DeleteStoreItem :exec
DELETE FROM store_items WHERE id = $1;

-- name: ListStoreItemsByStore :many
SELECT id, store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order 
FROM store_items WHERE store_id = $1 ORDER BY store_traversal_order ASC;

-- name: GetStoreItemsWithDetails :many
SELECT 
  si.id, 
  si.store_id, 
  s.name as store_name,
  si.item_id, 
  i.name as item_name,
  si.purchase_unit_id, 
  pu.name as purchase_unit_name,
  si.inventory_unit_id, 
  iu.name as inventory_unit_name,
  si.inventory_units_per_purchase, 
  si.purchased_by_decimal, 
  si.is_favorite, 
  si.store_traversal_order
FROM store_items si
JOIN stores s ON si.store_id = s.id
JOIN items i ON si.item_id = i.id
LEFT JOIN units pu ON si.purchase_unit_id = pu.id
LEFT JOIN units iu ON si.inventory_unit_id = iu.id
ORDER BY si.store_id ASC, si.store_traversal_order ASC LIMIT $1 OFFSET $2;

-- name: GetStoreItemsWithDetailsCount :one
SELECT COUNT(*) as count FROM store_items;

-- name: GetStoreItemsByStoreWithDetails :many
SELECT 
  si.id, 
  si.store_id, 
  si.item_id, 
  i.name as item_name,
  si.purchase_unit_id, 
  pu.name as purchase_unit_name,
  si.inventory_unit_id, 
  iu.name as inventory_unit_name,
  si.inventory_units_per_purchase, 
  si.purchased_by_decimal, 
  si.is_favorite, 
  si.store_traversal_order
FROM store_items si
JOIN items i ON si.item_id = i.id
LEFT JOIN units pu ON si.purchase_unit_id = pu.id
LEFT JOIN units iu ON si.inventory_unit_id = iu.id
WHERE si.store_id = $1
ORDER BY si.store_traversal_order ASC;

-- name: SearchStoreItemsByStore :many
SELECT id, store_id, item_id, purchase_unit_id, inventory_unit_id, inventory_units_per_purchase, purchased_by_decimal, is_favorite, store_traversal_order 
FROM store_items 
WHERE store_id = $1 AND item_id IN (SELECT id FROM items WHERE name ILIKE $2)
ORDER BY store_traversal_order ASC;
