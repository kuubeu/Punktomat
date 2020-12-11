package main

import (
	"fmt"
	"os"
	"punktomat/controller"
	"punktomat/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/scienceMagazine", controller.GetScienceMagazines)
}

func main() {
	app := fiber.New()

	if os.Getenv("ENVIRONMENT") == "DEV" {
		app.Use(cors.New())

	}

	initDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	setupRoutes(app)

	app.Listen(":4000")
}
