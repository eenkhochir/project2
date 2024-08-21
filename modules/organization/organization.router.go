package organization

import "github.com/gofiber/fiber/v2"

func OrganizationRoutes(app *fiber.App) {
	handler := OrganizationHandler{}

	app.Get("/organization", handler.OrganizationList)
	app.Get("/organization/:id", handler.FindOrganization)
	app.Post("/organizations", handler.Createorganization)
	app.Put("/organizations/:id", handler.Updateorganization)
	app.Delete("/organizations/:id", handler.Deleteorganization)
}
