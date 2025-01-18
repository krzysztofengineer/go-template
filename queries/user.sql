-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?
LIMIT
    1;