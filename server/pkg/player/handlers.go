package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Create(c *fiber.Ctx) error {
	var player *Player
	var err error
	var db *sql.DB
	var query string
	var statement *sql.Stmt

	player = new(Player)
	c.BodyParser(player)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	if err != nil {
		return c.Status(500).JSON(err)
	}

	query = `INSERT INTO players (name, password) VALUES (?,?);`
	statement, err = db.Prepare(query)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	_, err = statement.Exec(player.Name, player.Password)
	if err != nil {
		if sqlerror, ok := err.(*sqlite3.Error); ok {
			if sqlerror.Code == sqlite3.ErrConstraint {
				return c.Status(400).JSON(sqlerror.Code)
			}
			return c.Status(400).JSON(sqlerror.Code)
		}
		c.Status(400).JSON(err)
	}

	return c.JSON(player.Name)
}

func Get(c *fiber.Ctx) error {
	var db *sql.DB
	var err error
	var rows *sql.Rows
	var player string
	var players []string

	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	rows, err = db.Query(`SELECT name FROM players;`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&player)
		players = append(players, player)
	}

	return c.JSON(players)
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a player")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a player")
}
