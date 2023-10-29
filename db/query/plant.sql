-- name: CreatePlant :one
INSERT INTO plants (
 user_id,
 plant_name,
 species
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPlantFromId :one
SELECT * FROM plants
WHERE id = $1 LIMIT 1;
