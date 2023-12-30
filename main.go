package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"todoapp/handlers"
	"todoapp/middleware"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:        engine,
		ViewsLayout:  "layouts/base",
		ErrorHandler: handlers.HandleErrors,
	})

	app.Use(middleware.Compress)
	app.Use(middleware.Helmet)
	app.Use(middleware.BasicAuth)
	app.Use(middleware.Limiter)
	app.Static("/", "./public")

	// app.Get("/", handlers.HandleIndexGet)
	app.Get("/", handlers.HandleTodosGet)
	app.Get("/metrics", middleware.Monitor)

	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleTodosGet)
	todos.Post("/", handlers.HandleTodosPost)
	todos.Delete("/:id", handlers.HandleTodoDelete)
	todos.Get("/edit/:id", handlers.HandleTodoEditGet)
	todos.Patch("/update/:id", handlers.HandleTodoEditPatch)
	todos.Patch("/toggle/:id", handlers.HandleTodoTogglePatch)

	app.Listen(":3000")
}
