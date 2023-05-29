package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/muling3/auth/handlers"
)

func Router() *fiber.App {
	app := fiber.New()

	// Initializing default cors config
	app.Use(cors.New())

	// authenticating user
	app.Post("/", handlers.AuthenticateUser)

	// adding a user
	app.Post("/", handlers.CreateUser)

	//logging user out
	app.Post("/logout", handlers.LogoutUser)

	//test server
	app.Get("/test", handlers.TestServer)

	return app

}
