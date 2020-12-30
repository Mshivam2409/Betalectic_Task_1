package main

import (
	"betalectic_task_1/controller"
	"betalectic_task_1/database"

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
	// SignUp Route (Unauthenticated)
	app.Post("/signup", controller.SignUp)
	app.Post("/signin", controller.SignIn)
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// Restricted Route
	app.Get("/restricted", controller.GetDetails)
	app.Listen(":5050")
}
