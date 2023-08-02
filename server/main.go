package main

import (
	"log"
	"os"

	"github.com/NickDeChip/todo-list/database"
	routes "github.com/NickDeChip/todo-list/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.Connect("./todo.db")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000, https://localhost:6969",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("COOKIE_KEY"),
	}))

	app.Get("/", Root)

	routes.RegisterRoutes(app)

	app.Listen("localhost:6969")
}

func Root(c *fiber.Ctx) error {
	return c.SendString("Hey baby")
}
