package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"gateway-service/pkg/handler/health"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
)

func Health(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	healthCtrl := health.Controller{Validate: validate, Logger: logger}

	authUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.Post("/health", healthCtrl.Post)
		authUrl.Get("/health", healthCtrl.Get)
	}
	return app
}
