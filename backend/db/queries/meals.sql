-- name: GetMeal :one
SELECT id, name, meal_category_id, servings FROM meals WHERE id = $1;

-- name: GetMealByName :one
SELECT id, name, meal_category_id, servings FROM meals WHERE name = $1;

-- name: ListMeals :many
SELECT id, name, meal_category_id, servings FROM meals ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: ListMealsCount :one
SELECT COUNT(*) as count FROM meals;

-- name: CreateMeal :one
INSERT INTO meals (name, meal_category_id, servings) VALUES ($1, $2, $3) RETURNING id, name, meal_category_id, servings;

-- name: UpdateMeal :one
UPDATE meals SET name = $1, meal_category_id = $2, servings = $3 WHERE id = $4 RETURNING id, name, meal_category_id, servings;

-- name: DeleteMeal :exec
DELETE FROM meals WHERE id = $1;

-- name: ListMealsByCategory :many
SELECT id, name, meal_category_id, servings FROM meals WHERE meal_category_id = $1 ORDER BY name ASC LIMIT $2 OFFSET $3;

-- name: ListMealsByCategoryCount :one
SELECT COUNT(*) as count FROM meals WHERE meal_category_id = $1;

-- name: SearchMeals :many
SELECT id, name, meal_category_id, servings FROM meals WHERE name ILIKE $1 ORDER BY name ASC LIMIT $2 OFFSET $3;

-- name: SearchMealsCount :one
SELECT COUNT(*) as count FROM meals WHERE name ILIKE $1;

-- name: GetMealWithCategory :one
SELECT m.id, m.name, m.meal_category_id, mc.name as category_name, m.servings 
FROM meals m
JOIN meal_categories mc ON m.meal_category_id = mc.id
WHERE m.id = $1;
