package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		data := map[string]interface{}{
			"message": "Hello, World!",
		}

		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
