package main

import (
	"myapp/modules/organization"
	"myapp/modules/user"
	"myapp/modules/vehicle"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	user.UserRoutes(app)
	vehicle.VehicleRoutes(app)
	organization.OrganizationRoutes(app)
}
