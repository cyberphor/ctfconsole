package campaign

import (
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
)

type Campaign struct {
	Id    int         `json:"id"`
	Name  string      `json:"Campaign"`
	Teams []team.Team `json:"teams"`
}

func Post(c *fiber.Ctx) error {
	return c.SendString("Creating a campaign")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a campaign")
}

func Put(c *fiber.Ctx) error {
	return c.SendString("Updating a campaign")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a campaign")
}
