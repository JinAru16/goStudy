package main

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("welcome to Daily code buffer!")
}

func Routers(app *fiber.App) {
	app.Get("/user", user.GetUsers)
	app.Get("/user/:id", user.GetUsers)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUSer)
	app.Put("/user/:id", user.UpdateUSer)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")
}
