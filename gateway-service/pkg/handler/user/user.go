package user

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"gateway-service/utility"
)

func (base *Controller) CreateUser(c *fiber.Ctx) error {

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) Login(c *fiber.Ctx) error {

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
