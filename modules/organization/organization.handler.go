package organization

import (
	"myapp/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrganizationHandler struct{}

func (*OrganizationHandler) OrganizationList(c *fiber.Ctx) error {
	var organizations []Organization
	database.DB.Find(&organizations)
	return c.JSON(organizations)
}

func (*OrganizationHandler) FindOrganization(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	org := FindOrganization(uint(id))
	return c.JSON(org)
}

func (*OrganizationHandler) Createorganization(c *fiber.Ctx) error {
	//pointer
	organization := new(Organization)
	if err := c.BodyParser(organization); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&organization)
	return c.Status(201).JSON(organization)
}

func (*OrganizationHandler) Updateorganization(c *fiber.Ctx) error {
	id := c.Params("id")
	var organization Organization
	if result := database.DB.First(&organization, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "organization not found"})
	}
	if err := c.BodyParser(&organization); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&organization)
	return c.JSON(organization)
}

func (*OrganizationHandler) Deleteorganization(c *fiber.Ctx) error {
	id := c.Params("id")
	var organization Organization
	if result := database.DB.First(&organization, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "organization not found"})
	}
	database.DB.Delete(&organization)
	return c.SendStatus(204)
}
