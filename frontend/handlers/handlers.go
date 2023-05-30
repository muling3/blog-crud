package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muling3/frontend/utils"
)

func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"Title": "Hello Index",
	})
}

func RegisterPage(ctx *fiber.Ctx) error {
	return ctx.Render("register", fiber.Map{
		"Title": "Register Page",
	})
}

func FormHandler(ctx *fiber.Ctx) error {
	var user utils.User

	if err := ctx.BodyParser(&user); err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusBadRequest,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(errorMap)
	}

	data, _ := json.Marshal(user)

	bodyReader := bytes.NewReader(data)

	//switching on the action types:: "authenticate" | "create"
	switch user.Action {
	case "authenticate":
		return utils.AuthenticateUser(ctx, bodyReader)
	case "create":
		return utils.CreateUser(ctx, bodyReader)
	default:
		errorMap := fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "Unexpected action",
		}

		//return the error
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(errorMap)
	}
}
