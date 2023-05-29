package main

import (

	// "encoding/json"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/muling3/frontend/utils"
	// "github.com/muling3/frontend/utils"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {

	//engine
	engine := html.New("./views", ".html")

	//inityialising app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello Index",
		})
	})

	//register endpoint
	app.Get("/register", func(ctx *fiber.Ctx) error {
		log.Printf("%+v", ctx.Body())

		return ctx.Render("register", fiber.Map{
			"Title": "Register Page",
		})
	})

	//handling form actions
	app.Post("/", func(ctx *fiber.Ctx) error {
		var user utils.User

		if err := ctx.BodyParser(&user); err != nil {
			log.Println("error parsing", err)
		}

		data, _ := json.Marshal(user)

		bodyReader := bytes.NewReader(data)

		//sending data to auth service
		// 1. creating a http service
		switch user.Action {
		case "authenticate":
			request, err := http.NewRequest(http.MethodPost, "http://localhost:5000/", bodyReader)
			if err != nil {
				return ctx.SendStatus(http.StatusInternalServerError)
			}

			// 2. create an http client and run the request
			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return ctx.SendStatus(http.StatusInternalServerError)
			}

			defer response.Body.Close()

			resBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				os.Exit(1)
			}
			log.Printf("client: response body: %s\n", resBody)
		case "create":
			request, err := http.NewRequest(http.MethodPost, "http://localhost:5000/new", bodyReader)
			if err != nil {
				return ctx.SendStatus(http.StatusInternalServerError)
			}

			// 2. create an http client and run the request
			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return ctx.SendStatus(http.StatusInternalServerError)
			}

			defer response.Body.Close()

			resBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				os.Exit(1)
			}
			log.Printf("client: response body: %s\n", resBody)
		default:
			log.Println("Invalid errror")
		}

		return ctx.SendStatus(http.StatusCreated)
	})

	//serving static files
	app.Static("/", "./public")

	log.Fatal(app.Listen(":4100"))

}

func sendRequest() {

}
