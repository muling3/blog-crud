package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muling3/auth/utils"
)

func AuthenticateUser(ctx *fiber.Ctx) error {
	log.Println("authenticate action")
	var user utils.User

	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusBadRequest,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(errorMap)
	}
	log.Printf("%+v", user)

	responseMap := fiber.Map{
		"status":  http.StatusOK,
		"message": "Successfully hit server",
	}

	return ctx.JSON(responseMap)
}

func CreateUser(ctx *fiber.Ctx) error {
	var user utils.User

	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusBadRequest,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(errorMap)
	}
	
	// create user
	

	responseMap := fiber.Map{
		"status":  http.StatusCreated,
		"message": "Successfully created user",
	}
	return ctx.JSON(responseMap)
}

func LogoutUser(ctx *fiber.Ctx) error {
	responseMap := fiber.Map{
		"status":  http.StatusOK,
		"message": "Successfully logged out",
	}
	return ctx.JSON(responseMap)
}

func TestServer(ctx *fiber.Ctx) error {
	responseMap := fiber.Map{
		"status":  http.StatusAccepted,
		"message": "Server is up and running",
	}
	return ctx.JSON(responseMap)
}
