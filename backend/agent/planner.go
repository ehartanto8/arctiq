package agent

import "strings"

// Structured steps, after getting OpenAI response
type Planner struct{}

type Task struct {
	Step        int
	Instruction string
}

func (p *Planner) Plan(prompt string) ([]Task, error) {
	planText, err := getPlanFromGPT(prompt)
	if err != nil {
		return nil, err
	}
	/*
		1. Set up the project directory with Go modules and folders.
		2. Create a REST API in Go using Fiber.
		3. Add an endpoint to accept user prompts.
		4. Create a Next.js frontend with a form.
		5. Connect frontend to backend.
	*/
	lines := strings.Split(planText, "\n")
	var tasks []Task
	step := 1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Remove numeric from the beginning
		parts := strings.SplitN(line, ".", 2)
		if len(parts) < 2 {
			continue
		}

		instruction := strings.TrimSpace(parts[1])
		tasks = append(tasks, Task{Step: step, Instruction: instruction})
		step++
	}

	return tasks, nil
}
