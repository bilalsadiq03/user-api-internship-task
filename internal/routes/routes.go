package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bilalsadiq03/user-api-internship-task/internal/handler"
)

func Register(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Get("/users", userHandler.ListUsers)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Put("/users/:id", userHandler.UpdateUser)
}
