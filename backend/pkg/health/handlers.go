package health

import "github.com/gofiber/fiber/v2"

type Health struct {
	Status string
}

func Get(c *fiber.Ctx) error {
	health := new(Health)
	health.Status = "imok"
	return c.Status(200).JSON(health)
}
