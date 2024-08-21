package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	app.Get("/users", GetUsers)
	app.Get("/users/:id", GetUser)
	app.Post("/users", CreateUser)
	app.Put("/users/:id", UpdateUser)
	app.Delete("/users/:id", DeleteUser)
}
