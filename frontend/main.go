package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	//engine
	engine := html.New("./views", ".html")

	//inityialising app
	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello Index",
		})
	})

	//serving static files
	app.Static("/", "./public")

	log.Fatal(app.Listen(":4000"))

}