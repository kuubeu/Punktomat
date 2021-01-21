package controller

import (
	"punktomat/database"
	"punktomat/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Categories     []string `json:"categories"`
	Order          string   `json:"order"`
	OrderDirection string   `json:"orderDirection"`
	MinPoints      uint     `json:"minPoints"`
	MaxPoints      uint     `json:"maxPoints"`
}

// GetScienceMagazines returns all science magazines
func GetScienceMagazines(c *fiber.Ctx) error {
	db := database.DBConn
	var scienceMagazine []model.ScienceMagazine
	var count int64

	reqBody := new(Request)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	limit, err := strconv.Atoi(c.Query("limit", "100"))
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	search := "%" + c.Query("search", "") + "%"

	if err != nil {
		return c.Status(400).SendString("Invalid query param")
	}

	chain := db

	if reqBody.Order != "" {
		order := reqBody.Order

		if reqBody.OrderDirection != "" {
			order += " " + reqBody.OrderDirection
		}

		chain = chain.Order(order)
	}

	if reqBody.MinPoints != 0 {
		chain = chain.Where("points >= ?", reqBody.MinPoints)
	}

	if reqBody.MaxPoints != 0 {
		chain = chain.Where("points <= ?", reqBody.MaxPoints)
	}

	for _, category := range reqBody.Categories {
		chain = chain.Where(
			"categories @> ARRAY[?]", category)
	}

	chain = chain.Where(
		"title iLIKE ? OR second_title iLIKE ?", search, search).Offset(
		offset).Limit(limit)

	chain.Find(&scienceMagazine).Count(&count)

	return c.JSON(fiber.Map{
		"total": count, "results": scienceMagazine,
	})
}
