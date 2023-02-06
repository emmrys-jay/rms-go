package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"gateway-service/pkg/handler/auth"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
)

// Auth registers routes for authentication service
func Auth(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	authCtrl := auth.Controller{Validate: validate, Logger: logger}

	authUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.Post("/login", authCtrl.Login)
		authUrl.Post("/logout", authCtrl.Logout)
	}

	// Health routes for authentication service
	health := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		health.Post("/auth/health", authCtrl.Post)
		health.Get("/auth/health", authCtrl.Get)
	}
	return app
}
