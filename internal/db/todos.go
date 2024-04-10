package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"database/sql"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID    uuid.UUID
	Title string
	Done  bool
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0012"
	dbname   = "todo-db"
)

func getConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(5)

	if err != nil {
		return db, err
	}

	return db, nil
}

type Filter int

const (
	AllTodos Filter = iota
	DoneTodos
	NotDoneTodos
)

func GetTodos(filter Filter) ([]Todo, error) {
	var todos []Todo
	var err error
	var db *sql.DB
	db, err = getConnection()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if filter == DoneTodos {
		rows, err = db.Query("SELECT id, name, done FROM todos WHERE done ORDER BY pk")
	} else if filter == NotDoneTodos {
		rows, err = db.Query("SELECT id, name, done FROM todos WHERE NOT done ORDER BY pk")
	} else {
		rows, err = db.Query("SELECT id, name, done FROM todos ORDER BY pk")
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return todos, err
	}
	return todos, nil
}

func GetRemainngCount() (int, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}

	var count int
	err = db.QueryRow("SELECT count(*) FROM todos WHERE done = false").Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

// TODO: make errors visible to the user
func AddTodo(title string) (Todo, error) {
	var err error
	if title == "" {
		err = errors.New("The todo title cannot be empty!")
	}
	if err != nil {
		return Todo{}, err
	}
	db, err := getConnection()
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = db.QueryRow("INSERT INTO todos(id, name, done) VALUES ($1, $2, $3) RETURNING id, name, done", uuid.New(), title, false).Scan(&todo.ID, &todo.Title, &todo.Done)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(id uuid.UUID) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	var testID uuid.UUID
	err = db.QueryRow("DELETE FROM todos WHERE id = $1 RETURNING id", id).Scan(&testID)
	if err != nil {
		return err
	}

	if testID != id {
		return fmt.Errorf("Could not delete todo with ID: %s", id.String())
	}

	return nil
}

func GetTodo(id uuid.UUID) (Todo, error) {
	db, err := getConnection()
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = db.QueryRow("SELECT id, name, done FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Done)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func EditTodo(id uuid.UUID, title string) (Todo, error) {
	db, err := getConnection()
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = db.QueryRow("UPDATE todos SET name = $1 WHERE id = $2 RETURNING id, name, done", title, id).Scan(&todo.ID, &todo.Title, &todo.Done)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func ToggleTodo(id uuid.UUID) (Todo, error) {
	db, err := getConnection()
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = db.QueryRow("UPDATE todos SET done = NOT done WHERE id = $1 RETURNING id, name, done", id).Scan(&todo.ID, &todo.Title, &todo.Done)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
