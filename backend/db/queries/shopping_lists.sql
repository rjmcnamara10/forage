-- name: GetShoppingList :one
SELECT id, name, created_at FROM shopping_lists WHERE id = $1;

-- name: ListShoppingLists :many
SELECT id, name, created_at FROM shopping_lists ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ListShoppingListsCount :one
SELECT COUNT(*) as count FROM shopping_lists;

-- name: CreateShoppingList :one
INSERT INTO shopping_lists (name, created_at) VALUES ($1, CURRENT_TIMESTAMP) RETURNING id, name, created_at;

-- name: UpdateShoppingList :exec
UPDATE shopping_lists SET name = $1 WHERE id = $2;

-- name: DeleteShoppingList :exec
DELETE FROM shopping_lists WHERE id = $1;

-- name: ListShoppingListsByMostRecent :many
SELECT id, name, created_at FROM shopping_lists ORDER BY created_at DESC LIMIT $1;

-- name: GetShoppingListsContainingStoreItem :many
SELECT DISTINCT sl.id, sl.name, sl.created_at FROM shopping_lists sl
JOIN shopping_list_items sli ON sl.id = sli.shopping_list_id
WHERE sli.store_item_id = $1
ORDER BY sl.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetShoppingListsContainingStoreItemCount :one
SELECT COUNT(DISTINCT sl.id) as count FROM shopping_lists sl
JOIN shopping_list_items sli ON sl.id = sli.shopping_list_id
WHERE sli.store_item_id = $1;
