// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todo.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addTodo = `-- name: AddTodo :one
INSERT INTO todos(id, name, done) VALUES ($1, $2, $3) RETURNING id, name, done
`

type AddTodoParams struct {
	ID   uuid.UUID
	Name string
	Done bool
}

type AddTodoRow struct {
	ID   uuid.UUID
	Name string
	Done bool
}

func (q *Queries) AddTodo(ctx context.Context, arg AddTodoParams) (AddTodoRow, error) {
	row := q.db.QueryRowContext(ctx, addTodo, arg.ID, arg.Name, arg.Done)
	var i AddTodoRow
	err := row.Scan(&i.ID, &i.Name, &i.Done)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :one
DELETE FROM todos WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteTodo(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, deleteTodo, id)
	err := row.Scan(&id)
	return id, err
}

const editTodo = `-- name: EditTodo :one
UPDATE todos SET name = $1 WHERE id = $2 RETURNING id, name, done
`

type EditTodoParams struct {
	Name string
	ID   uuid.UUID
}

type EditTodoRow struct {
	ID   uuid.UUID
	Name string
	Done bool
}

func (q *Queries) EditTodo(ctx context.Context, arg EditTodoParams) (EditTodoRow, error) {
	row := q.db.QueryRowContext(ctx, editTodo, arg.Name, arg.ID)
	var i EditTodoRow
	err := row.Scan(&i.ID, &i.Name, &i.Done)
	return i, err
}

const getRemainingCount = `-- name: GetRemainingCount :one
SELECT count(*) FROM todos WHERE done = false
`

func (q *Queries) GetRemainingCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRemainingCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTodo = `-- name: GetTodo :one
SELECT name, id, done, pk, pos FROM todos WHERE id = $1
`

func (q *Queries) GetTodo(ctx context.Context, id uuid.UUID) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.Name,
		&i.ID,
		&i.Done,
		&i.Pk,
		&i.Pos,
	)
	return i, err
}

const listDoneTodos = `-- name: ListDoneTodos :many
SELECT name, id, done, pk, pos
FROM todos 
WHERE done 
ORDER BY pk
`

func (q *Queries) ListDoneTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listDoneTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.Name,
			&i.ID,
			&i.Done,
			&i.Pk,
			&i.Pos,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNotDoneTodos = `-- name: ListNotDoneTodos :many
SELECT name, id, done, pk, pos
FROM todos 
WHERE NOT done 
ORDER BY pk
`

func (q *Queries) ListNotDoneTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listNotDoneTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.Name,
			&i.ID,
			&i.Done,
			&i.Pk,
			&i.Pos,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTodos = `-- name: ListTodos :many
SELECT name, id, done, pk, pos
FROM todos 
ORDER BY pk
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.Name,
			&i.ID,
			&i.Done,
			&i.Pk,
			&i.Pos,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const toggleTodo = `-- name: ToggleTodo :one
UPDATE todos SET done = NOT done WHERE id = $1 RETURNING id, name, done
`

type ToggleTodoRow struct {
	ID   uuid.UUID
	Name string
	Done bool
}

func (q *Queries) ToggleTodo(ctx context.Context, id uuid.UUID) (ToggleTodoRow, error) {
	row := q.db.QueryRowContext(ctx, toggleTodo, id)
	var i ToggleTodoRow
	err := row.Scan(&i.ID, &i.Name, &i.Done)
	return i, err
}