package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func Create(c *fiber.Ctx) error {
	return c.SendString("Creating a player")
}

func Get(c *fiber.Ctx) error {
	var db *sql.DB
	var err error
	var rows *sql.Rows
	var id int
	var username string
	var player Player
	var players []Player

	db, err = sql.Open("sqlite3", "ctfconsole.db")
	rows, err = db.Query(`SELECT username FROM players;`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&id, &username)
		player = Player{
			Id:   id,
			Name: username,
		}
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
