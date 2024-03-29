package team

import (
	"github.com/gofiber/fiber/v2"
)

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"team"`
}

func Post(c *fiber.Ctx) error {
	return c.SendString("Creating a team")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a team")
}

func Put(c *fiber.Ctx) error {
	return c.SendString("Updating a team")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a team")
}
