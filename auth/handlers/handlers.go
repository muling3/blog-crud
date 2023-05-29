package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(ctx *fiber.Ctx) error {
	responseMap := fiber.Map{
		"status": http.StatusOK,
		"message": "Successfully hit server",
	}
	return ctx.JSON(responseMap)
}

func CreateUser(ctx *fiber.Ctx) error {
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
