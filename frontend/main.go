package main

import (
	"log"

	"github.com/muling3/frontend/routes"
)

func main() {
	// initialising app
	app := routes.Router()

	log.Fatal(app.Listen(":4100"))
}
