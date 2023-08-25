package handlers

import "github.com/gofiber/fiber/v2"

func CreateTeam(c *fiber.Ctx) error {
	return c.SendString("Creating a team")
}

func GetTeam(c *fiber.Ctx) error {
	return c.SendString("Getting a team")
}

func GetTeams(c *fiber.Ctx) error {
	return c.SendString("Getting all teams")
}

func UpdateTeam(c *fiber.Ctx) error {
	return c.SendString("Updating a team")
}

func DeleteTeam(c *fiber.Ctx) error {
	return c.SendString("Deleting a team")
}
