-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at;

-- name: CreateTodo :one
INSERT INTO todos (
  title, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;