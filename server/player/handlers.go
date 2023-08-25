package player

import "github.com/gofiber/fiber/v2"

func Create(c *fiber.Ctx) error {
	return c.SendString("Creating a player")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting all players")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a player")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a player")
}
