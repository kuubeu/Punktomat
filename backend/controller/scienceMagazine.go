package controller

import (
	"punktomat/database"
	"punktomat/model"

	"github.com/gofiber/fiber/v2"
)

// GetScienceMagazines returns all sience magazines
func GetScienceMagazines(c *fiber.Ctx) error {
	db := database.DBConn
	var scienceMagazine []model.ScienceMagazine
	db.Find(&scienceMagazine)
	return c.JSON(scienceMagazine)
}
