package controller

import (
	"punktomat/database"
	"punktomat/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetScienceMagazines returns all science magazines
func GetScienceMagazines(c *fiber.Ctx) error {
	db := database.DBConn
	var scienceMagazine []model.ScienceMagazine
	var count int64

	limit, err := strconv.Atoi(c.Query("limit", "100"))
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	search := "%" + c.Query("search", "") + "%"
	category := c.Query("category", "")

	if err != nil {
		return c.Status(400).SendString("Invalid query param")
	}

	chain := db

	if category != "" {
		chain = chain.Where(
			"categories @> ARRAY[?]", category)
	}

	db.Model(&model.ScienceMagazine{}).Count(&count)
	chain.Where(
		"title iLIKE ? OR second_title iLIKE ?", search, search).Offset(
		offset).Limit(limit).Find(
		&scienceMagazine)

	return c.JSON(fiber.Map{
		"total": count, "results": scienceMagazine,
	})
}
