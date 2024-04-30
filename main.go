package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"todoapp/internal/handlers"
	apiHandlers "todoapp/internal/handlers/api"
	"todoapp/internal/middleware"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/base",
	})

	app.Use(middleware.Compress)
	app.Use(middleware.Helmet)
	app.Use(middleware.BasicAuth)
	app.Use(middleware.Limiter)
	app.Static("/", "./public")

	app.Get("/", handlers.HandleTodosGet)
	app.Get("/metrics", middleware.Monitor)

	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleTodosGet)
	todos.Post("/", handlers.HandleTodosPost)
	todos.Delete("/:id", handlers.HandleTodoDelete)
	todos.Get("/edit/:id", handlers.HandleTodoEditGet)
	todos.Patch("/update/:id", handlers.HandleTodoEditPatch)
	todos.Patch("/toggle/:id", handlers.HandleTodoTogglePatch)
	todos.Get("/count", handlers.HandleTodosCountGet)
	todos.Post("/reorder", handlers.HandleTodosReorderPost)

	todoApi := app.Group("/api/todos")
	todoApi.Get("/", apiHandlers.HandleTodosGet)
	todoApi.Post("/", apiHandlers.HandleTodosPost)
	todoApi.Delete("/:id", apiHandlers.HandleTodoDelete)
	todoApi.Get("/edit/:id", apiHandlers.HandleTodoEditGet)
	todoApi.Patch("/update/:id", apiHandlers.HandleTodoEditPatch)
	todoApi.Patch("/toggle/:id", apiHandlers.HandleTodoTogglePatch)
	todoApi.Get("/count", apiHandlers.HandleTodosCountGet)
	todoApi.Post("/reorder", apiHandlers.HandleTodosReorderPost)

	app.Listen(":3000")
}
