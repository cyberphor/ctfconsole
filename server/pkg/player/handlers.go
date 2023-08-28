package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
)

type Player struct {
	Id       *int    `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func Create(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var db *sql.DB
	var query string
	var statement *sql.Stmt
	var sqlerror sqlite3.Error
	var ok bool

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	query = `INSERT INTO players (name, password) VALUES (?,?);`
	statement, err = db.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(err)
	}

	_, err = statement.Exec(player.Name, player.Password)
	if err != nil {
		sqlerror, ok = err.(sqlite3.Error)
		if ok {
			if sqlerror.Code == sqlite3.ErrConstraint {
				message["data"] = "player name is already taken"
				return c.Status(400).JSON(message)
			}
			message["data"] = err.Error()
			return c.Status(400).JSON(message)
		}
	}

	message["data"] = "created player"
	return c.Status(200).JSON(message)
}

func Get(c *fiber.Ctx) error {
	var message map[string][]Player
	var db *sql.DB
	var err error
	var rows *sql.Rows
	var player Player
	var players []Player

	message = make(map[string][]Player)
	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	if err != nil {
		var message map[string]string
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	rows, err = db.Query(`SELECT id, name FROM players;`)
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

func Update(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var db *sql.DB
	var query string
	var statement *sql.Stmt
	var sqlerror sqlite3.Error
	var ok bool

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	query = `UPDATE players SET name = (?) WHERE id = (?);`
	statement, err = db.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	_, err = statement.Exec(player.Name, player.Id)
	if err != nil {
		sqlerror, ok = err.(sqlite3.Error)
		if ok {
			message["data"] = sqlerror.Error()
			return c.Status(400).JSON(message)
		}
	}

	message["data"] = "updated player"
	return c.Status(200).JSON(message)
}

func Delete(c *fiber.Ctx) error {
	var message map[string]string
	var player *Player
	var err error
	var db *sql.DB
	var query string
	var statement *sql.Stmt
	var ok bool
	var sqlerror sqlite3.Error

	message = make(map[string]string)
	player = new(Player)
	c.BodyParser(player)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(400).JSON(message)
	}

	db, err = sql.Open("sqlite3", "storage/ctfconsole.db")
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(message)
	}

	query = `DELETE FROM players WHERE name = (?);`
	statement, err = db.Prepare(query)
	if err != nil {
		message["data"] = err.Error()
		return c.Status(500).JSON(err)
	}

	_, err = statement.Exec(player.Name)
	if err != nil {
		_, ok = err.(sqlite3.Error)
		if ok {
			message["data"] = sqlerror.Error()
			return c.Status(400).JSON(message)
		}
	}

	message["data"] = "deleted player"
	return c.Status(200).JSON(message)
}
