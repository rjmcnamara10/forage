-- name: GetShoppingListItem :one
SELECT id, shopping_list_id, store_item_id, purchase_quantity, note, custom_item_name, shopping_list_order 
FROM shopping_list_items WHERE id = $1;

-- name: ListShoppingListItems :many
SELECT id, shopping_list_id, store_item_id, purchase_quantity, note, custom_item_name, shopping_list_order 
FROM shopping_list_items WHERE shopping_list_id = $1 ORDER BY shopping_list_order ASC;

-- name: CreateShoppingListItem :one
INSERT INTO shopping_list_items (shopping_list_id, store_item_id, purchase_quantity, note, custom_item_name, shopping_list_order) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING id, shopping_list_id, store_item_id, purchase_quantity, note, custom_item_name, shopping_list_order;

-- name: UpdateShoppingListItem :exec
UPDATE shopping_list_items 
SET store_item_id = $1, purchase_quantity = $2, note = $3, custom_item_name = $4, shopping_list_order = $5 
WHERE id = $6;

-- name: DeleteShoppingListItem :exec
DELETE FROM shopping_list_items WHERE id = $1;

-- name: DeleteShoppingListItemsByShoppingListId :exec
DELETE FROM shopping_list_items WHERE shopping_list_id = $1;

-- name: GetShoppingListItemsWithDetails :many
SELECT 
  sli.id, 
  sli.shopping_list_id, 
  sli.store_item_id, 
  i.name as item_name,
  s.name as store_name,
  si.store_traversal_order,
  sli.purchase_quantity, 
  pu.name as purchase_unit_name,
  sli.note, 
  sli.custom_item_name,
  sli.shopping_list_order
FROM shopping_list_items sli
LEFT JOIN store_items si ON sli.store_item_id = si.id
LEFT JOIN items i ON si.item_id = i.id
LEFT JOIN stores s ON si.store_id = s.id
LEFT JOIN units pu ON si.purchase_unit_id = pu.id
WHERE sli.shopping_list_id = $1
ORDER BY s.name ASC, si.store_traversal_order ASC, sli.shopping_list_order ASC;

-- name: GetShoppingListItemsWithDetailsSimple :many
SELECT 
  sli.id, 
  sli.shopping_list_id, 
  sli.store_item_id, 
  sli.purchase_quantity, 
  sli.note, 
  sli.custom_item_name,
  sli.shopping_list_order
FROM shopping_list_items sli
WHERE sli.shopping_list_id = $1
ORDER BY sli.shopping_list_order ASC;

-- name: GetMaxShoppingListOrder :one
SELECT COALESCE(MAX(shopping_list_order), 0) as max_order FROM shopping_list_items WHERE shopping_list_id = $1;

-- name: UpdateShoppingListItemsOrder :exec
UPDATE shopping_list_items SET shopping_list_order = $1 WHERE id = $2;
