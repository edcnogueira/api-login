package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("🥉 Last handler")
		return c.JSON("Hello, World 👋!")
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Created!")
		return ctx.JSON("Created")
	})

	err := app.Listen(":3000")

	if err != nil {
		return
	}
}
