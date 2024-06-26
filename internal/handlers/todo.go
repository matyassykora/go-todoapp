package handlers

import (
	"time"
	"todoapp/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var demoSleepTime = 1 * time.Second

func HandleTodosGet(c *fiber.Ctx) error {
	filter := c.Query("filter")
	var todos db.Todos
	var err error

	if filter != "" {
		time.Sleep(demoSleepTime)
	}

	switch filter {
	case "done":
		todos, err = db.GetTodos(db.DoneTodos)

	case "notdone":
		todos, err = db.GetTodos(db.NotDoneTodos)

	// Assuming filter is 'all' or empty
	default:
		todos, err = db.GetTodos(db.AllTodos)
		filter = "all"
	}

	if err != nil {
		return err
	}

	var count int
	count, err = db.GetRemainingCount()

	if err != nil {
		return err
	}

	return c.Render("todos", fiber.Map{
		"Todos":  todos,
		"Filter": filter,
		"Count":  count,
	})
}

func HandleTodosCountGet(c *fiber.Ctx) error {
	var count int
	var err error
	count, err = db.GetRemainingCount()

	if err != nil {
		return err
	}

	return c.Render("partials/todos/count", fiber.Map{
		"Count": count,
	}, "")
}

func HandleTodosPost(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	title := c.FormValue("title")
	todo, err := db.AddTodo(title)
	if err != nil {
		return err
	}

	c.Response().Header.Add("Hx-Trigger", "updateCount")

	return c.Render("partials/todos/todo", fiber.Map{
		"ID":    todo.ID,
		"Title": todo.Title,
		"Done":  todo.Done,
	}, "")
}

func HandleTodoDelete(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	err = db.DeleteTodo(id)
	if err != nil {
		return err
	}

	c.Response().Header.Add("Hx-Trigger", "updateCount")

	return nil
}

func HandleTodoEditGet(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)

	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}

	todo, err := db.GetTodo(id)
	if err != nil {
		return err
	}

	return c.Render("partials/todos/edit-todo", fiber.Map{
		"ID":    todo.ID,
		"Title": todo.Title,
		"Done":  todo.Done,
	}, "")
}

func HandleTodoEditPatch(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	title := c.FormValue("title")
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}
	todo, err := db.EditTodo(id, title)
	if err != nil {
		return err
	}

	return c.Render("partials/todos/todo", fiber.Map{
		"ID":    todo.ID,
		"Title": todo.Title,
		"Done":  todo.Done,
	}, "")
}

func HandleTodoTogglePatch(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}

	todo, err := db.ToggleTodo(id)
	if err != nil {
		return err
	}

	c.Response().Header.Add("Hx-Trigger", "updateCount")

	return c.Render("partials/todos/todo", fiber.Map{
		"ID":    todo.ID,
		"Title": todo.Title,
		"Done":  todo.Done,
	}, "")
}

// TODO: this is very inefficient - does it matter?
func HandleTodosReorderPost(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	filter := c.Query("filter")
	var err error
	var dbTodos db.Todos

	if filter == "done" {
		dbTodos, err = db.GetTodos(db.DoneTodos)
	} else if filter == "notdone" {
		dbTodos, err = db.GetTodos(db.NotDoneTodos)
	} else {
		// assuming filter is 'all' or doesn't exist
		dbTodos, err = db.GetTodos(db.AllTodos)
		filter = "all"
	}

	if err != nil {
		return err
	}

	todoPostArgs := c.Request().PostArgs().PeekMulti("todo")
	todos := db.Todos{}
	for _, todo := range todoPostArgs {
		id, err := uuid.Parse(string(todo))
		if err != nil {
			return err
		}

		dbTodo, err := dbTodos.Find(id)
		if err != nil {
			return err
		}

		todos = append(todos, db.Todo{
			ID:    id,
			Title: dbTodo.Title,
			Done:  dbTodo.Done,
		})
	}

	for i := 0; i < len(todos); i++ {
		err = db.ReorderTodo(&todos[i])
		if err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"Todos": todos,
	})
}
