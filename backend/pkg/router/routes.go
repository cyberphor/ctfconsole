package router

import (
	"database/sql"

	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/scoreboard"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Status string `default:"imok"`
}

func (h Health) Get(c *fiber.Ctx) error {
	var message map[string]string

	message = make(map[string]string)
	message["data"] = h.Status
	return c.Status(200).JSON(message)
}

func Route(app *fiber.App, db *sql.DB) {
	var health *Health
	var ph *player.Handler

	// get health handler
	health = &Health{}

	// get player handler
	ph = &player.Handler{
		DB: db,
	}

	// set routes for api health data
	app.Get("/api/v1/ruok", health.Get)

	// set routes for player data
	app.Post("/api/v1/player", ph.Create)
	app.Get("/api/v1/player", ph.Get)
	app.Put("/api/v1/player", ph.Update)
	app.Delete("/api/v1/player", ph.Delete)

	// set routes for admin data
	app.Get("/api/v1/admin", admin.Get)
	app.Post("/api/v1/admin", admin.Update)

	// set routes for team data
	app.Get("/team", team.Get)
	app.Post("/team", team.Update)

	// set routes for challenge data
	app.Get("/challenge", challenge.Get)
	app.Post("/challenge", challenge.Update)

	// set routes for scoreboard data
	app.Get("/scoreboard", scoreboard.Get)
	app.Post("/scoreboard", scoreboard.Update)
}
