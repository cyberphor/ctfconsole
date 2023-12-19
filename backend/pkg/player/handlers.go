package player

import (
	"github.com/cyberphor/ctfconsole/pkg/database"
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
	db, err := database.Connect()
	if err != nil {
		// return Internal Server Error
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// prepare SQL statement
	query := `INSERT INTO players (name, password) VALUES ($1, $2);`
	statement, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// execute SQL statement
	result, err := statement.Exec(player.Name, player.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.Status(200).JSON(result)
}

func Get(c *fiber.Ctx) error {
	// parse GET request
	player := new(Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// connect to database
	db, err := database.Connect()
	if err != nil {
		// return Internal Server Error
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// prepare SQL statement
	query := `SELECT * FROM players WHERE name = ($1);`
	statement, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	// execute SQL statement
	result, err := statement.Exec(player.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.Status(200).JSON(result)
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
