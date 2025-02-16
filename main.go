package main

import (
	db "demo/db"
	"demo/routes"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("error while connecting to database: ", err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal("error while migrating to database: ", err)
	}
	app := fiber.New()
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Demo server is working fine",
		})
	})
	routes.SetupRoutes(app)
	err = app.Listen(":8000")
	if err != nil {
		log.Println("error while starting the server: ", err)
	}
}
