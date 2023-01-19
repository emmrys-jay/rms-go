package main

import (
	"github.com/go-playground/validator/v10"
	"log"

	"gateway-service/internal/config"
	"gateway-service/utility"

	// "github.com/workshopapps/pictureminer.api/pkg/repository/storage/redis"
	"gateway-service/pkg/router"
)

func init() {
	config.Setup()
	// redis.SetupRedis() uncomment when you need redis

}

func main() {
	//Load config
	logger := utility.NewLogger()
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	app := router.Setup(validatorRef, logger)

	logger.Info("Server is starting at 127.0.0.1:%s", getConfig.Server.Port)
	log.Fatal(app.Listen(":" + getConfig.Server.Port))
}
