package main

import (
	"database/sql"

	"github.com/cyberphor/ctfconsole/admin"
	"github.com/cyberphor/ctfconsole/challenge"
	"github.com/cyberphor/ctfconsole/player"
	"github.com/cyberphor/ctfconsole/scoreboard"
	"github.com/cyberphor/ctfconsole/team"
	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type Team struct {
	Id      int      `json:"id"`
	Name    string   `json:"team"`
	Players []Player `json:"players"`
}

type Challenge struct {
	Id       int    `json:"id"`
	Name     string `json:"challenge"`
	Points   int    `json:"points"`
	Solution string `json:"solution"`
}

type Scoreboard struct {
	Id    int    `json:"id"`
	Name  string `json:"scoreboard"`
	Teams []Team `json:"teams"`
}

func setQueries() []string {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS players (
		id INTEGER PRIMARY KEY,
		username  TEXT,
		password  TEXT,
		UNIQUE(username)
		);`,
	}
	return queries
}

func setTable(database *sql.DB, query string) {
	// define and execute a query to create the Admin table
	statement, err := database.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func setDatabase(name string) {
	// if it doesn't already exist, create the database
	ctfconsoleDatabase, err := sql.Open("sqlite3", name)
	if err != nil {
		panic(err)
	}

	for _, query := range setQueries() {
		setTable(ctfconsoleDatabase, query)
	}
	ctfconsoleDatabase.Close()
}

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setRoutes(app *fiber.App) {
	// default route
	app.Get("/", Index)

	// player routes
	app.Get("/api/v1/player", player.Get)
	app.Get("/api/v1/player/{playerId}", player.Get)
	app.Post("/api/v1/player", player.Update)

	// admin routes
	app.Get("/api/v1/admin", admin.Get)
	app.Get("/api/v1/admin/{adminId}", admin.Get)
	app.Post("/api/v1/admin", admin.Update)

	// team routes
	app.Get("/api/v1/team", team.Get)
	app.Get("/api/v1/team/{teamId}", team.Get)
	app.Post("/api/v1/team", team.Update)

	// challenge routes
	app.Get("/api/v1/challenge", challenge.Get)
	app.Get("/api/v1/challenge/{challengeId}", challenge.Get)
	app.Post("/api/v1/challenge", challenge.Update)

	// scoreboard routes
	app.Get("/api/v1/scoreboard", scoreboard.Get)
	app.Get("/api/v1/scoreboard/{scoreboardId}", scoreboard.Get)
	app.Post("/api/v1/scoreboard", scoreboard.Update)
}

func main() {
	setDatabase("ctfconsole.db")
	var app *fiber.App = fiber.New()
	setRoutes(app)
	app.Listen(":9001")
}
