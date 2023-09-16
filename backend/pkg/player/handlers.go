package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       *int    `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

type Handler struct {
	DB *sql.DB
}

func (h Handler) Create(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var query string
	var statement *sql.Stmt

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	query = `INSERT INTO players (name, password) VALUES (?,?);`
	statement, err = h.DB.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(err)
	}

	_, err = statement.Exec(player.Name, player.Password)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	message["data"] = "created player"
	return c.Status(200).JSON(message)
}

func (h Handler) Get(c *fiber.Ctx) error {
	var message map[string][]Player
	var err error
	var rows *sql.Rows
	var player Player
	var players []Player

	message = make(map[string][]Player)
	rows, err = h.DB.Query(`SELECT id, name FROM players;`)
	if err != nil {
		var message map[string]string
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	for rows.Next() {
		rows.Scan(&player.Id, &player.Name)
		players = append(players, player)
	}

	message["data"] = players
	return c.Status(200).JSON(message)
}

func (h Handler) Update(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var query string
	var statement *sql.Stmt

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	query = `UPDATE players SET name = (?) WHERE id = (?);`
	statement, err = h.DB.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	_, err = statement.Exec(player.Name, player.Id)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	message["data"] = "updated player"
	return c.Status(200).JSON(message)
}

func (h Handler) Delete(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var query string
	var statement *sql.Stmt

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	query = `DELETE FROM players WHERE name = (?);`
	statement, err = h.DB.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(err)
	}

	_, err = statement.Exec(player.Name)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	message["data"] = "deleted player"
	return c.Status(200).JSON(message)
}
