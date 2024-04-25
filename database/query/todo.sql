-- name: ListTodos :many
SELECT *
FROM todos 
ORDER BY pk;

-- name: ListDoneTodos :many
SELECT *
FROM todos 
WHERE done 
ORDER BY pk;

-- name: ListNotDoneTodos :many
SELECT *
FROM todos 
WHERE NOT done 
ORDER BY pk;

-- name: AddTodo :one
INSERT INTO todos(id, name, done) VALUES ($1, $2, $3) RETURNING id, name, done;

-- name: GetRemainingCount :one
SELECT count(*) FROM todos WHERE done = false;

-- name: DeleteTodo :one
DELETE FROM todos WHERE id = $1 RETURNING id;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: EditTodo :one
UPDATE todos SET name = $1 WHERE id = $2 RETURNING id, name, done;

-- name: ToggleTodo :one
UPDATE todos SET done = NOT done WHERE id = $1 RETURNING id, name, done;
