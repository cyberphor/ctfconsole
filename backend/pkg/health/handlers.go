package health

import "github.com/gofiber/fiber/v2"

type Handler struct {
	Status string
}

func (h Handler) Get(c *fiber.Ctx) error {
	h.Status = "imok"
	message := make(map[string]string)
	message["data"] = h.Status
	return c.Status(200).JSON(message)
}
