package controllers

import (
	"goDemoApi/app/models"
	"goDemoApi/database"

	"github.com/gofiber/fiber/v2"
)

func UsersShow(c *fiber.Ctx) error {
	var user models.User

	database.Instance().First(&user, c.Params("id"))

	return c.JSON(user)
}
