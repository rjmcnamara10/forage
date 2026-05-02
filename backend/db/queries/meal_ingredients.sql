-- name: GetMealIngredient :one
SELECT id, meal_id, ingredient_id, unit_id, required_amount, is_seasoning, depletes_slowly, alternate_ingredient_id 
FROM meal_ingredients WHERE id = $1;

-- name: ListMealIngredients :many
SELECT id, meal_id, ingredient_id, unit_id, required_amount, is_seasoning, depletes_slowly, alternate_ingredient_id 
FROM meal_ingredients WHERE meal_id = $1 ORDER BY id ASC;

-- name: CreateMealIngredient :one
INSERT INTO meal_ingredients (meal_id, ingredient_id, unit_id, required_amount, is_seasoning, depletes_slowly, alternate_ingredient_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING id, meal_id, ingredient_id, unit_id, required_amount, is_seasoning, depletes_slowly, alternate_ingredient_id;

-- name: UpdateMealIngredient :one
UPDATE meal_ingredients 
SET ingredient_id = $1, unit_id = $2, required_amount = $3, is_seasoning = $4, depletes_slowly = $5, alternate_ingredient_id = $6 
WHERE id = $7 RETURNING id, meal_id, ingredient_id, unit_id, required_amount, is_seasoning, depletes_slowly, alternate_ingredient_id;

-- name: DeleteMealIngredient :exec
DELETE FROM meal_ingredients WHERE id = $1;

-- name: DeleteMealIngredientsByMealId :exec
DELETE FROM meal_ingredients WHERE meal_id = $1;

-- name: GetMealIngredientsWithDetails :many
SELECT 
  mi.id, 
  mi.meal_id, 
  mi.ingredient_id, 
  i.name as ingredient_name,
  mi.unit_id, 
  u.name as unit_name,
  mi.required_amount, 
  mi.is_seasoning, 
  mi.depletes_slowly, 
  mi.alternate_ingredient_id,
  ai.name as alternate_ingredient_name
FROM meal_ingredients mi
JOIN items i ON mi.ingredient_id = i.id
LEFT JOIN units u ON mi.unit_id = u.id
LEFT JOIN items ai ON mi.alternate_ingredient_id = ai.id
WHERE mi.meal_id = $1
ORDER BY mi.id ASC;

-- name: ListMealIngredientsForMultipleMeals :many
SELECT 
  mi.meal_id,
  mi.ingredient_id,
  mi.unit_id,
  mi.required_amount,
  mi.alternate_ingredient_id
FROM meal_ingredients mi
WHERE mi.meal_id = ANY($1::int[]);
