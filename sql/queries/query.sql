-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
where ID = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);


-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ?
where id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = ?;