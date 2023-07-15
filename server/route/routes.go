package routes

import (
	"github.com/NickDeChip/todo-list/route/todo"
	"github.com/NickDeChip/todo-list/route/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	todo.RegisterRoutes(app)
	user.RedirectRoutes(app)
}
