-- name: CreateUser :exec
INSERT INTO TEST_USER (
    name,
    phone,
    birthday,
    password
) VALUES (
             $1, $2, $3, $4
         );