package controllers

import (
	"goDemoApi/app/models"
	"goDemoApi/database"

	"github.com/gofiber/fiber/v2"
)

type NotFoundResponse struct {
	Errors  map[string][]string
	Message string
}

func UsersShow(c *fiber.Ctx) error {
	var user models.User

	result := database.Instance().First(&user, c.Params("id"))

	if result.RowsAffected == 0 {
		errors := make(map[string][]string)
		errors["NotFound"] = []string{"User not found."}

		return c.Status(fiber.StatusNotFound).JSON(&NotFoundResponse{
			Errors:  errors,
			Message: "There was a problem processing your request",
		})
	}

	return c.JSON(user)
}
