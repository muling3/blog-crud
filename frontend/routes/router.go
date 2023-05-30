package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/muling3/frontend/handlers"
)

func Router() *fiber.App {
	//engine
	engine := html.New("./views", ".html")

	//inityialising app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", handlers.LoginPage)

	//register endpoint
	app.Get("/register", handlers.RegisterPage)

	//handling form actions
	app.Post("/", handlers.FormHandler)

	//serving static files
	app.Static("/", "./public")

	return app
}
