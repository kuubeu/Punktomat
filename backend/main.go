package main

import (
	"fmt"
	"strconv"

	"os"
	"punktomat/controller"
	"punktomat/database"
	"punktomat/model"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"))

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successful")

	database.DBConn.AutoMigrate(&model.ScienceMagazine{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/scienceMagazine", controller.GetScienceMagazines)
}

func initModel(filename string, labelsRow int) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		panic("Failed to open .xlsx file")
	}
	fmt.Println(".xlsx file opened")

	rows, err := file.GetRows("Czasopisma")

	for i := labelsRow + 1 ; i < len(rows); i++ {
		var categories []int64
		for index, colCell := range rows[i] {	
			if colCell == "x" {
				number,_ := strconv.ParseInt(rows[labelsRow][index], 10, 64)
				categories = append(categories, number)
			}
		}
		points, _ := strconv.Atoi(rows[i][7])
		magazine := model.ScienceMagazine{
			Title: rows[i][1],
			Issn: rows[i][2],
			Eissn: rows[i][3],
			SecondTitle: rows[i][4],
			SecondIssn: rows[i][5],
			SecondEissn: rows[i][6],
			Points: points,
			Categories: pq.Int64Array(categories),
		}
		database.DBConn.Create(&magazine)
	}
}

func main() {
	app := fiber.New()

	if os.Getenv("ENVIRONMENT") == "DEV" {
	 	app.Use(cors.New())
	}

	initDatabase()
	
	var sa []model.ScienceMagazine
	if err := database.DBConn.Where("id = ?", "1").First(&sa).Error; err != nil {
		initModel("./wykaz.xlsx", 3)
		app.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(sa)
		})
	}

	setupRoutes(app)

	app.Listen(":4000")
}
