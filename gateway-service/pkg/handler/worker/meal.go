package worker

import (
	"gateway-service/utility"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (base *Controller) CreateMeal(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetMeal(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetAllMeals(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) UpdateMeal(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) DeleteMeal(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
