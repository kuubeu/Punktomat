package controller

import (
	"punktomat/database"
	"punktomat/model"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Categories     []string `json:"categories"`
	Order          string   `json:"order"`
	OrderDirection string   `json:"orderDirection"`
	MinPoints      uint     `json:"minPoints"`
	MaxPoints      uint     `json:"maxPoints"`
	Limit          int      `json:"limit"`
	Offset         int      `json:"offset"`
	Search         string   `json:"search"`
}

// GetScienceMagazines returns all science magazines
func GetScienceMagazines(c *fiber.Ctx) error {
	db := database.DBConn
	var scienceMagazine []model.ScienceMagazine
	var count int64
	var limit int

	reqBody := new(Request)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if reqBody.Limit == 0 {
		limit = 100
	} else {
		limit = reqBody.Limit
	}

	search := "%" + reqBody.Search + "%"

	chain := db

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

	chain = chain.Where("title iLIKE ? OR second_title iLIKE ?", search, search).Limit(limit)

	if reqBody.Offset != 0 {
		chain = chain.Offset(reqBody.Offset)
	}

	if reqBody.Order != "" {
		order := "upper(" + reqBody.Order + ")"

		if reqBody.OrderDirection != "" {
			order += " " + reqBody.OrderDirection
		}

		chain = chain.Order(order)
	} else {
		chain = chain.Order("id")
	}

	chain.Find(&scienceMagazine).Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"total": count, "results": scienceMagazine,
	})
}
