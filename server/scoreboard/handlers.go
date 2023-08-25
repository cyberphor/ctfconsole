package scoreboard

import "github.com/gofiber/fiber/v2"

func Create(c *fiber.Ctx) error {
	return c.SendString("Creating a scoreboard")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a scoreboard")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a scoreboard")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a scoreboard")
}
