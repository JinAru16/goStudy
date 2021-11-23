package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	micro := fiber.New()
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app := fiber.New()
	app.Mount("/john", micro) // GET /john/doe -> 200 OK
	log.Fatal(app.Listen(":3000"))
}
