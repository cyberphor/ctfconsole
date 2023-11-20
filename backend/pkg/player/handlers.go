package player

import (
	"fmt"

	"github.com/cyberphor/ctfconsole/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Id       *int    `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (h Handler) Create(c *fiber.Ctx) error {
	db, err := config.GetDatabaseConnection()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	db.Query("")
	fmt.Println("Username: %s, Password: %s", h.Name, h.Password)
	return c.Status(200).JSON(h.Name)
}
