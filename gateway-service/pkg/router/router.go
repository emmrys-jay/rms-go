package router

import (
	"gateway-service/pkg/middleware"
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(validate *validator.Validate, logger *utility.Logger) *fiber.App {
	app := fiber.New()

	// Middlewares
	// r.Use(gin.Logger())
	app.Use(middleware.Logger())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowOrigins:     "*",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	ApiVersion := "v1"
	Health(app, validate, ApiVersion, logger)
	Auth(app, validate, ApiVersion, logger)

	return app
}
