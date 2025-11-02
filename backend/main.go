package main

import (
	"arctiq-backend/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

// Workflow
// Main -> Agent (Route) -> Planner (Logic) -> OpenAI -> Planner -> Frontend

// Server
func main() {
	app := fiber.New()

	// Recover
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) error {
		id := c.Get("X-Request-ID")
		if id == "" {
			id = fmt.Sprintf("%x", c.Context().ID())
		}
		c.Set("X-Request-ID", id)
		c.Locals("reqid", id)
		return c.Next()
	})

	// Log
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${latency} ${method} ${path} reqid=${locals:reqid}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Arctiq backend is running! Use POST /agent to test.")
	})

	app.Post("/agent", routes.AgentHandler)

	log.Println("Arctiq backend listening on: 8080")
	log.Fatal(app.Listen(":8080"))
}
