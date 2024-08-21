package vehicle

import (
	"myapp/database"

	"github.com/gofiber/fiber/v2"
)

func Getvehicles(c *fiber.Ctx) error {
	var vehicles []Vehicle
	db := database.DB
	db.Find(&vehicles)
	return c.JSON(vehicles)
}

func Getvehicle(c *fiber.Ctx) error {
	id := c.Params("id")
	var vehicle Vehicle
	db := database.DB
	if err := db.First(&vehicle, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "vehicle not found"})
	}
	return c.JSON(vehicle)
}

func Createvehicle(c *fiber.Ctx) error {
	//pointer
	vehicle := new(Vehicle)
	if err := c.BodyParser(vehicle); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	db := database.DB
	db.Create(&vehicle)
	return c.Status(201).JSON(vehicle)
}

func Updatevehicle(c *fiber.Ctx) error {
	id := c.Params("id")
	var vehicle Vehicle
	db := database.DB
	if result := db.First(&vehicle, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "vehicle not found"})
	}
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	db.Save(&vehicle)
	return c.JSON(vehicle)
}

func Deletevehicle(c *fiber.Ctx) error {
	id := c.Params("id")
	var vehicle Vehicle
	db := database.DB
	if err := db.First(&vehicle, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "vehicle not found"})
	}
	db.Delete(&vehicle)
	return c.SendStatus(204)
}
