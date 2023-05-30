package main

import (
	"context"
	"database/sql"
	"log"

	database "github.com/muling3/auth/db/sqlc"
	"github.com/muling3/auth/routes"

	_ "github.com/lib/pq"
)

func main() {
	// initializing app
	app := routes.Router()

	// initializing database
	ctx := context.Background()

	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/blog_microservice?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to database successfully")

	dbConn := database.New(db)

	// list all users
	users, err := dbConn.ListUsers(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println(users)

	//serve
	log.Fatal(app.Listen(":5000"))
}
