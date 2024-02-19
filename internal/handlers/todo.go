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
	var todos []db.Todo
	var err error

	if filter == "done" {
		todos, err = db.GetTodos(db.DoneTodos)
	} else if filter == "notdone" {
		todos, err = db.GetTodos(db.NotDoneTodos)
	} else {
		// assuming filter is 'all' or doesn't exist
		todos, err = db.GetTodos(db.AllTodos)
		filter = "all"
	}

	if err != nil {
		return err
	}

	var count int
	count, err = db.GetRemainngCount()

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
	count, err = db.GetRemainngCount()

	if err != nil {
		return err
	}

	return c.Render("partials/todos/count", fiber.Map{
		"Count": count,
	},"")
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
	id := c.Params("id")
	err := db.DeleteTodo(uuid.MustParse(id))
	if err != nil {
		return err
	}

	c.Response().Header.Add("Hx-Trigger", "updateCount")

	return nil
}

func HandleTodoEditGet(c *fiber.Ctx) error {
	time.Sleep(demoSleepTime)
	id := c.Params("id")
	todo, err := db.GetTodo(uuid.MustParse(id))
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
	id := c.Params("id")
	title := c.FormValue("title")
	todo, err := db.EditTodo(uuid.MustParse(id), title)
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
	id := c.Params("id")
	todo, err := db.ToggleTodo(uuid.MustParse(id))
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
