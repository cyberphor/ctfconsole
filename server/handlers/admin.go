package handlers

import "github.com/gofiber/fiber/v2"

func CreateAdmin(c *fiber.Ctx) error {
	return c.SendString("Creating an admin")
}

func GetAdmin(c *fiber.Ctx) error {
	return c.SendString("Getting an admin")
}

func GetAdmins(c *fiber.Ctx) error {
	return c.SendString("Getting all admins")
}

func UpdateAdmin(c *fiber.Ctx) error {
	return c.SendString("Updating an admin")
}

func DeleteAdmin(c *fiber.Ctx) error {
	return c.SendString("Deleting an admin")
}
