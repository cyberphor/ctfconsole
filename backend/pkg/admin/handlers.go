package admin

import "github.com/gofiber/fiber/v2"

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func Create(c *fiber.Ctx) error {
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
