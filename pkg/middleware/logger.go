package middleware

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger() fiber.Handler {

	// instantiation
	logger := logrus.New()

	//Set output
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{})

	//Set log level
	logger.SetLevel(logrus.DebugLevel)

	return func(c *fiber.Ctx) error {
		//Start time
		startTime := time.Now()

		//Process request
		c.Next()

		//End time
		endTime := time.Now()

		//Execution time
		latencyTime := endTime.Sub(startTime)

		//Request method
		reqMethod := c.Method()

		//Request routing
		reqUri := c.OriginalURL()

		// status code
		//statusCode := c.Writer.Status()

		// request IP
		clientIP := c.IP()

		//Log format
		logger.WithFields(logrus.Fields{
			//"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()

		return nil
	}
}
