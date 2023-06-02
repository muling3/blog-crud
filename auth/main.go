package main

import (
	"database/sql"
	"log"

	"github.com/muling3/auth/config"
	database "github.com/muling3/auth/db/sqlc"
	"github.com/muling3/auth/handlers"
	"github.com/muling3/auth/routes"

	_ "github.com/lib/pq"
)

func main() {
	// initialising root level config
	var appConfig config.AppConfig

	// initializing database
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/blog_microservice?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to database successfully")

	dbConn := database.New(db)

	//registering root config
	appConfig.Db = dbConn

	//initialising app config
	handlers.NewAppConfig(&appConfig)

	// initializing app
	app := routes.Router()

	//serve
	log.Fatal(app.Listen(":5000"))
}
