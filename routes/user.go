package routes

import (
	db "demo/db"
	"demo/models"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	if db.DB == nil {
		return c.Status(500).SendString("database connection is not initialized")
	}
	rows, err := db.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Println("error while running the query on the database: ", err)
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var users []models.Users
	for rows.Next() {
		var user models.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("error while scanning the rows: ", err)
			return c.Status(500).SendString(err.Error())
		}
		users = append(users, user)
	}
	for _, user := range users {
		log.Println("Name: " + user.Name)
		log.Println("Email: " + user.Email)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
    user := new(models.Users)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    err := db.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(200).JSON(user)
}

func SetupRoutes(app *fiber.App) {
	app.Get("/users", GetAllUsers)
	app.Post("/user", AddUser)
}
