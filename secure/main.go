package main

import (
	"betalectic_task_1_secure/controller"
	"betalectic_task_1_secure/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	app := fiber.New()
	database.IntializeMongoDB()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Secure Microservice is up!")
	})
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// Restricted Route
	app.Get("/restricted", controller.GetDetails)
	app.Listen(":5080")
}
