package todo

import (
	"net/http"

	"github.com/NickDeChip/todo-list/database"
	"github.com/NickDeChip/todo-list/middleware/auth"
	"github.com/NickDeChip/todo-list/model"
	"github.com/NickDeChip/todo-list/utility"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	todos := app.Group("/api/todo")

	todos.Use(auth.New())

	// C
	todos.Post("/", PostTodos)

	// R
	todos.Get("/all", GetTodos)

	// D
	todos.Delete("/", DeleteTodo)

}

func GetTodos(c *fiber.Ctx) error {
	userId, err := auth.GetUserID(c)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}
	todos, err := database.GetConnection().Query("SELECT ID, Info FROM Todo WHERE UserID=$1", userId)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}
	defer todos.Close()

	result := make([]model.Todo, 0)

	for todos.Next() {
		todo := model.Todo{}
		todos.Scan(&todo.ID, &todo.Info)
		result = append(result, todo)
	}

	return c.JSON(result)
}

func PostTodos(c *fiber.Ctx) error {
	userId, err := auth.GetUserID(c)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}

	todoPostData := model.Info{}
	if err := c.BodyParser(&todoPostData); err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	res, err := database.GetConnection().Exec("INSERT INTO Todo (Info, UserId) VALUES ($1, $2)", todoPostData.Info, userId)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	return c.JSON(model.ID{ID: id})
}

func DeleteTodo(c *fiber.Ctx) error {
	userId, err := auth.GetUserID(c)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}

	todoID := model.ID{}
	err = c.BodyParser(&todoID)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	res, err := database.GetConnection().Exec("DELETE FROM Todo WHERE Id=$1 AND UserID=$2", todoID.ID, userId)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	id, err := res.RowsAffected()
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	return c.JSON(model.ID{ID: id})
}
