package router

import (
	"fmt"
	"gateway-service/pkg/handler/worker"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Worker registers routes for worker service
func Worker(app *fiber.App, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *fiber.App {

	workerCtrl := worker.Controller{Validate: validate, Logger: logger}

	// Meal routes for worker service
	mealUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		mealUrl.Post("/meal", workerCtrl.CreateMeal)
		mealUrl.Get("/meal/:id", workerCtrl.GetMeal)
		mealUrl.Get("/meal/get_all", workerCtrl.GetAllMeals)
		mealUrl.Put("/meal/:id", workerCtrl.UpdateMeal)
		mealUrl.Delete("/meal/:id", workerCtrl.DeleteMeal)
	}

	// Table routes for worker service
	tableUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		tableUrl.Post("/table", workerCtrl.CreateTable)
		tableUrl.Get("/table/:id", workerCtrl.GetTable)
		tableUrl.Get("/table/get_all", workerCtrl.GetAllTables)
		tableUrl.Put("/table/:id", workerCtrl.UpdateTable)
		tableUrl.Delete("/table/:id", workerCtrl.DeleteTable)
	}

	// Menu routes for worker service
	menuUrl := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		menuUrl.Post("/menu", workerCtrl.CreateMenu)
		menuUrl.Get("/menu/:id", workerCtrl.GetMenu)
		menuUrl.Get("/menu/get_all", workerCtrl.GetAllMenus)
		menuUrl.Put("/menu/:id", workerCtrl.UpdateMenu)
		menuUrl.Delete("/menu/:id", workerCtrl.DeleteMenu)
	}

	// Health routes for worker service
	health := app.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		health.Post("/worker/health", workerCtrl.Post)
		health.Get("/worker/health", workerCtrl.Get)
	}

	return app
}
