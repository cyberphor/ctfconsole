package challenge

import "github.com/gofiber/fiber/v2"

type Challenge struct {
	Id       int    `json:"id"`
	Name     string `json:"challenge"`
	Points   int    `json:"points"`
	Solution string `json:"solution"`
}

func Create(c *fiber.Ctx) error {
	return c.SendString("Creating a challenge")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a challenge")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a challenge")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a challenge")
}
