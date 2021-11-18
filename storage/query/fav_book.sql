-- name: GetFavBook :one
SELECT *
FROM fav_book
WHERE user_id = $1 AND book_id = $2
LIMIT 1;

-- name: GetAllFavBooks :many
SELECT *
FROM fav_book
ORDER BY created_at desc OFFSET $1
LIMIT $2;

-- name: CreateFavBook :one
INSERT INTO fav_book (
        book_id,
        user_id
    )
VALUES ($1, $2)
RETURNING *;
-- name: UpdateFavBook :exec
UPDATE fav_book
SET book_id = $1
WHERE user_id = $2 AND book_id = $3;

-- name: DeleteFavBook :exec
DELETE FROM fav_book
WHERE book_id = $1 AND user_id = $2;

-- name: GetUsersFavBooks :many
SELECT fb.user_id, fb.book_id, b.name, b.image FROM fav_book as fb LEFT JOIN book as b on fb.book_id = b.id WHERE user_id = $1;