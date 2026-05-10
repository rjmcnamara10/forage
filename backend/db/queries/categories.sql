-- name: GetItemCategory :one
SELECT id, name FROM item_categories WHERE id = $1;

-- name: GetItemCategoryByName :one
SELECT id, name FROM item_categories WHERE name = $1;

-- name: ListItemCategories :many
SELECT id, name FROM item_categories ORDER BY name ASC;

-- name: ListItemCategoriesDesc :many
SELECT id, name FROM item_categories ORDER BY name DESC;

-- name: CreateItemCategory :one
INSERT INTO item_categories (name) VALUES ($1) RETURNING id, name;

-- name: UpdateItemCategory :one
UPDATE item_categories SET name = $1 WHERE id = $2 RETURNING id, name;

-- name: DeleteItemCategory :exec
DELETE FROM item_categories WHERE id = $1;

-- name: GetMealCategory :one
SELECT id, name FROM meal_categories WHERE id = $1;

-- name: GetMealCategoryByName :one
SELECT id, name FROM meal_categories WHERE name = $1;

-- name: ListMealCategories :many
SELECT id, name FROM meal_categories ORDER BY name ASC;

-- name: ListMealCategoriesDesc :many
SELECT id, name FROM meal_categories ORDER BY name DESC;

-- name: CreateMealCategory :one
INSERT INTO meal_categories (name) VALUES ($1) RETURNING id, name;

-- name: UpdateMealCategory :one
UPDATE meal_categories SET name = $1 WHERE id = $2 RETURNING id, name;

-- name: DeleteMealCategory :exec
DELETE FROM meal_categories WHERE id = $1;
