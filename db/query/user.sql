
-- name: CreateUser :exec
INSERT INTO TEST_USER (
    name,
    phone,
    birthday,
    password
) VALUES (?, ?, ?, ?);


-- name: GetUser :one
SELECT * FROM TEST_USER
WHERE name = ? LIMIT 1;