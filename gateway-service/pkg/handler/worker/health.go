package worker

import (
	"fmt"
	"gateway-service/internal/model"
	"gateway-service/service/ping"
	"gateway-service/utility"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (base *Controller) Post(c *fiber.Ctx) error {
	var (
		req = model.Ping{}
	)

	err := c.BodyParser(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		return c.JSON(rd)

	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		return c.JSON(rd)
	}
	if !ping.ReturnTrue() {
		rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", "ping failed", fmt.Errorf("ping failed"), nil)
		return c.JSON(rd)
	}
	base.Logger.Info("ping successful")

	rd := utility.BuildSuccessResponse(http.StatusOK, "ping successful", req.Message)
	return c.JSON(rd)

}

func (base *Controller) Get(c *fiber.Ctx) error {
	if !ping.ReturnTrue() {
		rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", "ping failed", fmt.Errorf("ping failed"), nil)
		return c.JSON(rd)
	}
	base.Logger.Info("ping successful")
	rd := utility.BuildSuccessResponse(http.StatusOK, "ping successful", "user object")
	return c.JSON(rd)
}
