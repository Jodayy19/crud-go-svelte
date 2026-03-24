package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 🔥 fungsi untuk cari ID berikutnya
func getNextID(users []User) int {
	maxID := 0
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	users := []User{
		{ID: 1, Name: "Jod"},
		{ID: 2, Name: "Budi"},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API jalan",
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, user := range users {
			userID := strconv.Itoa(user.ID)

			if id == userID {
				return c.JSON(user)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.JSON(fiber.Map{"message": "Deleted"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var input User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		if input.Name == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
		}

		for i, user := range users {
			if user.ID == id {
				users[i].Name = input.Name
				return c.JSON(users[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		newUser := new(User)

		if err := c.BodyParser(newUser); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		if newUser.Name == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Name wajib diisi",
			})
		}

		// 🔥 pakai auto ID dari data
		newUser.ID = getNextID(users)

		users = append(users, *newUser)

		return c.JSON(newUser)
	})

	app.Listen(":3000")
}