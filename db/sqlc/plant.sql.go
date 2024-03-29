// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: plant.sql

package db

import (
	"context"
	"database/sql"
)

const createPlant = `-- name: CreatePlant :one
INSERT INTO plants (
 user_id,
 plant_name,
 species
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, plant_name, species, created_at, updated_at
`

type CreatePlantParams struct {
	UserID    sql.NullInt64  `json:"user_id"`
	PlantName string         `json:"plant_name"`
	Species   sql.NullString `json:"species"`
}

func (q *Queries) CreatePlant(ctx context.Context, arg CreatePlantParams) (Plant, error) {
	row := q.db.QueryRowContext(ctx, createPlant, arg.UserID, arg.PlantName, arg.Species)
	var i Plant
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PlantName,
		&i.Species,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlantFromId = `-- name: GetPlantFromId :one
SELECT id, user_id, plant_name, species, created_at, updated_at FROM plants
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPlantFromId(ctx context.Context, id int64) (Plant, error) {
	row := q.db.QueryRowContext(ctx, getPlantFromId, id)
	var i Plant
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PlantName,
		&i.Species,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
