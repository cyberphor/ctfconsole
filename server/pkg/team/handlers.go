package team

import (
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/gofiber/fiber/v2"
)

type Team struct {
	Id      int             `json:"id"`
	Name    string          `json:"team"`
	Players []player.Player `json:"players"`
}

func Create(c *fiber.Ctx) error {
	return c.SendString("Creating a team")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a team")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a team")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a team")
}