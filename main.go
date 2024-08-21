package main

import (
	"log"
	"myapp/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()

	InitRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
