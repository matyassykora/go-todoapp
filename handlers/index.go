package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleIndexGet(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Name": "PEPA",
	})
}

