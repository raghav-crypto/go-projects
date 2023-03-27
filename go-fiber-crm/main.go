package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/raghav-crypto/go-projects/database"
	"github.com/raghav-crypto/go-projects/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead/:id", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = database.DBConn.AutoMigrate(&lead.Lead{})
	if err != nil {
		panic("failed to migrate database")

	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
