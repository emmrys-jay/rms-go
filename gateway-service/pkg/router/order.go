package router

import (
	"fmt"
	"gateway-service/pkg/handler/order"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Order registers routes for order service
func Order(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	orderCtrl := order.Controller{Validate: validate, Logger: logger}

	orderUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		orderUrl.Post("/order", orderCtrl.OrderMeal)
		orderUrl.Get("/order/:id", orderCtrl.GetOrder)
		orderUrl.Get("/order/get/:userID", orderCtrl.GetOrders)
	}

	// Health routes for order service
	health := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		health.Post("/order/health", orderCtrl.Post)
		health.Get("/order/health", orderCtrl.Get)
	}
	return app
}
