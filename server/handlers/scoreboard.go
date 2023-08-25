package handlers

import "github.com/gofiber/fiber/v2"

func CreateScoreboard(c *fiber.Ctx) error {
	return c.SendString("Creating a scoreboard")
}

func GetScoreboard(c *fiber.Ctx) error {
	return c.SendString("Getting a scoreboard")
}

func GetScoreboards(c *fiber.Ctx) error {
	return c.SendString("Getting all scoreboards")
}

func UpdateScoreboard(c *fiber.Ctx) error {
	return c.SendString("Updating a scoreboard")
}

func DeleteScoreboard(c *fiber.Ctx) error {
	return c.SendString("Deleting a scoreboard")
}
