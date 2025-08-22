package routes

import (
	"arctiq-backend/agent"

	"github.com/gofiber/fiber/v2"
)

// Request Handler
func AgentHandler(c *fiber.Ctx) error {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Planner
	planner := agent.Planner{}
	tasks, err := planner.Plan(req.Prompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(tasks)
}
