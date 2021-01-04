package main

import (
	"betalectic_task_1/controller"
	"betalectic_task_1/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	database.IntializeMongoDB()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Auth Microservice is up!")
	})
	// SignUp Route (Unauthenticated)
	app.Post("/signup", controller.SignUp)
	app.Post("/signin", controller.SignIn)
	app.Listen(":5050")
}
