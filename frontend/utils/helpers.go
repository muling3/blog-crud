package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(ctx *fiber.Ctx, body io.Reader) error {
	request, err := http.NewRequest(http.MethodPost, "http://localhost:5000/", body)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	// create an http client and run the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	//unmarshalling response body
	var res ResponseBody
	if err = json.Unmarshal(resBody, &res); err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	return ctx.JSON(fiber.Map{
		"status":  res.Status,
		"message": res.Message,
	})
}

func CreateUser(ctx *fiber.Ctx, body io.Reader) error {
	request, err := http.NewRequest(http.MethodPost, "http://localhost:5000/new", body)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	// 2. create an http client and run the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	//unmarshalling response body
	var res ResponseBody
	if err = json.Unmarshal(resBody, &res); err != nil {
		errorMap := fiber.Map{
			"error":   http.StatusInternalServerError,
			"message": err,
		}

		//return the error
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(errorMap)
	}

	return ctx.JSON(fiber.Map{
		"status":  res.Status,
		"message": res.Message,
	})
}
