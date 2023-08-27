package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Create(c *fiber.Ctx) error {
	var db *sql.DB
	var err error
	var query string
	var statement *sql.Stmt
	var player Player

	player = Player{
		Name:     c.Params("name"),
		Password: c.Params("password"),
	}

	db, err = sql.Open("sqlite3", "ctfconsole.db")
	if err != nil {
		panic(err)
	}

	query = `INSERT OR IGNORE INTO players (name, password) VALUES (?,?);`
	statement, err = db.Prepare(query)
	if err != nil {
		panic(err)
	}

	statement.Exec(player.Name, player.Password)
	if err != nil {
		panic(err)
	}

	return c.JSON(player)
}

func Get(c *fiber.Ctx) error {
	var db *sql.DB
	var err error
	var rows *sql.Rows
	var player Player
	var players []Player

	db, err = sql.Open("sqlite3", "ctfconsole.db")
	rows, err = db.Query(`SELECT id, name FROM players;`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&player.Id, &player.Name)
		players = append(players, player)
		if err != nil {
			return c.JSON(err)
		}
	}
	return c.JSON(players)
}

func Update(c *fiber.Ctx) error {
	return c.SendString("Updating a player")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Deleting a player")
}
