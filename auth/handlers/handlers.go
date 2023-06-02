package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muling3/auth/config"
	database "github.com/muling3/auth/db/sqlc"
	"github.com/muling3/auth/utils"
)

var appConfig *config.AppConfig

// initialising app config
func NewAppConfig(app *config.AppConfig) {
	appConfig = app
}

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
	// initialising db instance
	db := appConfig.Db

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
	args := database.CreateUserParams{
		FullName:    user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		AccountType: database.NullAccountType{AccountType: database.AccountTypeRegular, Valid: true},
	}

	usr, err := db.CreateUser(context.Background(), args)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	responseMap := fiber.Map{
		"status":  http.StatusCreated,
		"message": fmt.Sprintf("Successfully created %s with ID %d", usr.FullName, usr.ID),
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
	db := appConfig.Db

	//get all users
	users, err := db.ListUsers(context.Background())
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	responseMap := fiber.Map{
		"status":  http.StatusAccepted,
		"message": users,
	}
	return ctx.JSON(responseMap)
}
