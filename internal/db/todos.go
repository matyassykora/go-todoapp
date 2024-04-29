package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"database/sql"

	"todoapp/database/sqlc"

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
	ctx := context.Background()
	var todos []Todo
	var err error
	var db *sql.DB
	db, err = getConnection()
	if err != nil {
		return nil, err
	}

	queries := database.New(db)
	var rows []database.Todo

	if filter == DoneTodos {
		rows, err = queries.ListDoneTodos(ctx)
	} else if filter == NotDoneTodos {
		rows, err = queries.ListNotDoneTodos(ctx)
	} else {
		rows, err = queries.ListTodos(ctx)
	}

	if err != nil {
		return nil, err
	}

	for _, todo := range rows {
		todos = append(todos, Todo{
			ID:    todo.ID,
			Title: todo.Name,
			Done:  todo.Done,
		})
	}

	return todos, nil
}

func GetRemainingCount() (int, error) {
	ctx := context.Background()
	db, err := getConnection()
	if err != nil {
		return -1, err
	}

	queries := database.New(db)

	count, err := queries.GetRemainingCount(ctx)

	if err != nil {
		return -1, err
	}

	return int(count), nil
}

func AddTodo(title string) (*Todo, error) {
	ctx := context.Background()
	var err error
	if title == "" {
		err = errors.New("The todo title cannot be empty!")
	}
	if err != nil {
		return nil, fiber.NewError(400, err.Error())
	}
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	var row database.AddTodoRow
	row, err = queries.AddTodo(ctx, database.AddTodoParams{
		ID:   uuid.New(),
		Name: title,
		Done: false,
	})

	if err != nil {
		return nil, err
	}

	todo := &Todo{
		ID:    row.ID,
		Title: row.Name,
		Done:  row.Done,
	}

	return todo, nil
}

func DeleteTodo(id uuid.UUID) error {
	ctx := context.Background()
	db, err := getConnection()
	if err != nil {
		return err
	}

	queries := database.New(db)

	var testID uuid.UUID
	testID, err = queries.DeleteTodo(ctx, id)

	if err != nil {
		return err
	}

	if testID != id {
		return fmt.Errorf("Could not delete todo with ID: %s", id.String())
	}

	return nil
}

func GetTodo(id uuid.UUID) (*Todo, error) {
	ctx := context.Background()
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	var dbTodo database.Todo
	dbTodo, err = queries.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	todo := &Todo{
		ID:    dbTodo.ID,
		Title: dbTodo.Name,
		Done:  dbTodo.Done,
	}

	return todo, nil
}

func EditTodo(id uuid.UUID, title string) (*Todo, error) {
	ctx := context.Background()
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	if title == "" {
		err = errors.New("The todo title cannot be empty!")
	}
	if err != nil {
		return nil, fiber.NewError(400, err.Error())
	}

	queries := database.New(db)

	var row database.EditTodoRow
	row, err = queries.EditTodo(ctx, database.EditTodoParams{
		ID:   id,
		Name: title,
	})

	if err != nil {
		return nil, err
	}

	todo := &Todo{
		ID:    row.ID,
		Title: row.Name,
		Done:  row.Done,
	}
	return todo, nil
}

func ToggleTodo(id uuid.UUID) (*Todo, error) {
	ctx := context.Background()
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	var row database.ToggleTodoRow
	row, err = queries.ToggleTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	todo := &Todo{
		ID:    row.ID,
		Title: row.Name,
		Done:  row.Done,
	}

	return todo, nil
}
