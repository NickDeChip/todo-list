package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/mattn/go-sqlite3"
)

type todo struct {
	ID   int64  `json:"id"`
	Info string `json:"info"`
}

type ID struct {
	ID int64 `json:"id"`
}

type Info struct {
	Info string `json:"info"`
}

var db = connectToDB()

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", root)

	// C
	app.Post("/todo", postTodos)

	// R
	app.Get("/todos", getTodos)

	// D
	app.Delete("/todo", deleteTodo)

	app.Listen("localhost:6969")
}

func root(c *fiber.Ctx) error {
	return c.SendString("Hey baby")
}

func getTodos(c *fiber.Ctx) error {
	todos, err := db.Query("SELECT * FROM Todo")
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer todos.Close()

	result := make([]todo, 0)

	for todos.Next() {
		todo := todo{}
		todos.Scan(&todo.ID, &todo.Info)
		result = append(result, todo)
	}

	return c.JSON(result)
}

func postTodos(c *fiber.Ctx) error {
	todoPostData := Info{}
	if err := c.BodyParser(&todoPostData); err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	res, err := db.Exec("INSERT INTO Todo (Info) VALUES ($1)", todoPostData.Info)
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.JSON(ID{ID: id})
}

func deleteTodo(c *fiber.Ctx) error {
	todoID := ID{}
	err := c.BodyParser(&todoID)
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	res, err := db.Exec("DELETE FROM Todo WHERE Id=$1", todoID.ID)
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	id, err := res.RowsAffected()
	if err != nil {
		log.Printf("%s \n", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.JSON(ID{ID: id})
}

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
