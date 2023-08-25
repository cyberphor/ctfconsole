package admin

import "github.com/gofiber/fiber/v2"

func CreateAdmin(c *fiber.Ctx) error {
	return c.SendString("Creating an admin")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting an admin")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating an admin")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting an admin")
}
