-- name: GetBook :one
SELECT *
FROM book
WHERE id = $1
LIMIT 1;
-- name: GetBooks :many
SELECT *
FROM book 
WHERE (name ilike '%' || @search::varchar || '%' OR image ilike '%' || @search::varchar || '%')
OFFSET @_offset LIMIT @_limit;
-- name: CreateBook :one
INSERT INTO book (
        id,
        name,
        image
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateBook :exec
UPDATE book
SET name = $1,
    image = $2
WHERE id = $3;
-- name: DeleteBook :exec
DELETE FROM book
WHERE id = $1;
-- name: GetCount :one
SELECT count(1)
FROM book 
WHERE (name ilike '%' || $1::varchar || '%' OR image ilike '%' || $1::varchar || '%');