package worker

import (
	"gateway-service/utility"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (base *Controller) CreateTable(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetTable(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) GetAllTables(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) UpdateTable(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}

func (base *Controller) DeleteTable(c *fiber.Ctx) error {

	// TODO
	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", "")
	return c.JSON(rd)

}
