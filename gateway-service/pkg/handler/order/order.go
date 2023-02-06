package order

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"gateway-service/utility"
)

func (base *Controller) OrderMeal(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetOrder(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetOrders(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
