package main

import (
	"log"

	"github.com/muling3/auth/routes"
)

func main() {
	app := routes.Router()

	//serve
	log.Fatal(app.Listen(":5000"))
}
