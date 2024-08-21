package vehicle

import "github.com/gofiber/fiber/v2"

func VehicleRoutes(app *fiber.App) {
	app.Get("/vehicle", Getvehicles)
	app.Get("/vehicle/:id", Getvehicle)
	app.Post("/vehicles", Createvehicle)
	app.Put("/vehicles/:id", Updatevehicle)
	app.Delete("/vehicles/:id", Deletevehicle)
}
