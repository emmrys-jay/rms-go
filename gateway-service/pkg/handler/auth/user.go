package auth

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"gateway-service/utility"
)

func (base *Controller) Login(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) Logout(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
