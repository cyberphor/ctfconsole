package scoreboard

import (
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
)

type Scoreboard struct {
	Id    int         `json:"id"`
	Name  string      `json:"scoreboard"`
	Teams []team.Team `json:"teams"`
}

func Post(c *fiber.Ctx) error {
	return c.SendString("Creating a scoreboard")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Getting a scoreboard")
}

func Put(c *fiber.Ctx) error {
	return c.SendString("Updating a scoreboard")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a scoreboard")
}
