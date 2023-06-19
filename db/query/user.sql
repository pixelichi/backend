-- name: CreateUser :one
INSERT INTO users (
 username,
 hashed_password,
 email
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: GetUserFromId :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserFromUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
