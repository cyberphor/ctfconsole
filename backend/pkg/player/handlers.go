package player

import (
	"github.com/cyberphor/ctfconsole/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func Post(c *fiber.Ctx) error {
	// parse POST request
	player := new(Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// connect to database
	db, err := config.DatabaseConnection()
	if err != nil {
		// return Internal Server Error
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// prepare SQL statement
	query := `INSERT INTO players (name, password) VALUES (?,?);`
	statement, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// execute SQL statement
	_, err = statement.Exec(player.Name, player.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.Status(200).JSON("Created player")
}

func Get(c *fiber.Ctx) error {
	p := new(Player)
	c.BodyParser(p)
	return c.Status(200).JSON(p)
}

func Put(c *fiber.Ctx) error {
	p := new(Player)
	c.BodyParser(p)
	return c.Status(200).JSON(p)
}

func Delete(c *fiber.Ctx) error {
	p := new(Player)
	c.BodyParser(p)
	return c.Status(200).JSON(p)
}
