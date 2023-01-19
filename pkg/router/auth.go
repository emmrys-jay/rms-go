package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"gateway-service/pkg/handler/user"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
)

func Auth(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	auth := user.Controller{Validate: validate, Logger: logger}

	authUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.Post("/create_user", auth.CreateUser)
		authUrl.Post("/login", auth.Login)
	}
	return app
}
