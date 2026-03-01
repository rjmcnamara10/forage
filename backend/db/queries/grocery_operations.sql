-- name: GetMealIngredientsForIds :many
SELECT 
  mi.meal_id,
  mi.ingredient_id,
  mi.unit_id,
  mi.required_amount,
  mi.is_seasoning,
  mi.depletes_slowly,
  mi.alternate_ingredient_id
FROM meal_ingredients mi
WHERE mi.meal_id = ANY($1::int[])
ORDER BY mi.meal_id, mi.id;

-- name: GetInventoryForItems :many
SELECT 
  ii.item_id,
  ii.unit_id,
  ii.stored_amount
FROM inventory_items ii
WHERE ii.item_id = ANY($1::int[]);

-- name: GetMealServings :one
SELECT servings FROM meals WHERE id = $1;

-- name: BatchGetMeals :many
SELECT id, name, meal_category_id, servings FROM meals WHERE id = ANY($1::int[]);

-- name: UpdateInventoryBatch :exec
UPDATE inventory_items SET stored_amount = $1 WHERE item_id = $2;

-- name: GetStoreItemsForItems :many
SELECT 
  si.id,
  si.store_id,
  si.item_id,
  s.name as store_name,
  i.name as item_name,
  si.purchase_unit_id,
  si.inventory_unit_id,
  si.inventory_units_per_purchase,
  si.purchased_by_decimal,
  si.store_traversal_order,
  pu.name as purchase_unit_name,
  iu.name as inventory_unit_name
FROM store_items si
JOIN stores s ON si.store_id = s.id
JOIN items i ON si.item_id = i.id
LEFT JOIN units pu ON si.purchase_unit_id = pu.id
LEFT JOIN units iu ON si.inventory_unit_id = iu.id
WHERE si.item_id = ANY($1::int[])
ORDER BY s.id, si.store_traversal_order;
