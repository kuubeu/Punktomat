package main

import (
	"fmt"
	"strconv"

	"os"
	"punktomat/controller"
	"punktomat/database"
	"punktomat/enum"
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
	app.Post("/scienceMagazine", controller.GetScienceMagazines)
}

func initModel(filename string, labelsRow int) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		panic("Failed to open .xlsx file")
	}
	fmt.Println(".xlsx file opened")

	rows, err := file.GetRows("Czasopisma")

	for i := labelsRow + 1; i < len(rows); i++ {
		var categories []string
		for index, colCell := range rows[i] {
			if colCell == "x" {
				number, _ := strconv.ParseInt(rows[labelsRow][index], 10, 64)
				categories = append(categories, enum.Category[number])
			}
		}
		points, _ := strconv.Atoi(rows[i][8])
		magazine := model.ScienceMagazine{
			Title:       rows[i][2],
			Issn:        rows[i][3],
			Eissn:       rows[i][4],
			SecondTitle: rows[i][5],
			SecondIssn:  rows[i][6],
			SecondEissn: rows[i][7],
			Points:      points,
			Categories:  pq.StringArray(categories),
		}
		database.DBConn.Create(&magazine)
	}

	fmt.Println("Magazines successfully loaded")
}

func main() {
	app := fiber.New()

	if os.Getenv("ENVIRONMENT") == "DEV" {
		app.Use(cors.New())
	} else {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "https://punktomat.herokuapp.com",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	initDatabase()
	setupRoutes(app)

	var scienceMagazine []model.ScienceMagazine
	if err := database.DBConn.First(&scienceMagazine).Error; err != nil {
		initModel("./wykaz.xlsx", 1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	app.Listen(":" + port)

}
