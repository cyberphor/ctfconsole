package handlers

import "github.com/gofiber/fiber/v2"

func CreatePlayer(c *fiber.Ctx) error {
	return c.SendString("Creating a player")
}

func GetPlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(401).SendString("Invalid Player ID")
	}
	// for i, p in players, send back c.JSON(player)
	return c.SendString("Getting a player")
}

func GetPlayers(c *fiber.Ctx) error {
	return c.SendString("Getting all players")
}

func UpdatePlayer(c *fiber.Ctx) error {
	return c.SendString("Updating a player")
}

func DeletePlayer(c *fiber.Ctx) error {
	return c.SendString("Deleting a player")
}
