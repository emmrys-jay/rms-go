package router

import (
	"fmt"
	"gateway-service/pkg/handler/user"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// User registers routes for user service
func User(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	userCtrl := user.Controller{Validate: validate, Logger: logger}

	userUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		userUrl.Post("/user", userCtrl.CreateUser)
		userUrl.Get("/user/:id", userCtrl.GetUser)
		userUrl.Get("/user/get_all", userCtrl.GetAllUsers)
		userUrl.Put("/user/:id", userCtrl.UpdateUser)
		userUrl.Delete("/user/:id", userCtrl.DeleteUser)
	}

	// Health routes for user service
	health := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		health.Post("/user/health", userCtrl.Post)
		health.Get("/user/health", userCtrl.Get)
	}
	return app
}
