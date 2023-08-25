package handlers

import "github.com/gofiber/fiber/v2"

func CreateChallenge(c *fiber.Ctx) error {
	return c.SendString("Creating a challenge")
}

func GetChallenge(c *fiber.Ctx) error {
	return c.SendString("Getting a challenge")
}

func GetChallenges(c *fiber.Ctx) error {
	return c.SendString("Getting all challenges")
}

func UpdateChallenge(c *fiber.Ctx) error {
	return c.SendString("Updating a challenge")
}

func DeleteChallenge(c *fiber.Ctx) error {
	return c.SendString("Deleting a challenge")
}
