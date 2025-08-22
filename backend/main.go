package main

import (
	"arctiq-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/agent", routes.AgentHandler)

	log.Println("Arctiq backend listening on: 8080")
	log.Fatal(app.Listen(":8080"))
}
