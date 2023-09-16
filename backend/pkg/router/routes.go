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

type apiHealth struct {
	Status string `default:"imok"`
}

func (api apiHealth) Get(c *fiber.Ctx) error {
	var message map[string]string

	message = make(map[string]string)
	message["data"] = api.Status
	return c.Status(200).JSON(message)
}

func Route(app *fiber.App, db *sql.DB) {
	var r fiber.Router
	var api *apiHealth
	var ph *player.Handler

	// summarize route prefixes
	r = app.Group("/api/v1")

	// get health handler
	api = &apiHealth{}

	// get player handler
	ph = &player.Handler{
		DB: db,
	}

	// set routes for api health data
	app.Get("/ruok", api.Get)

	// set routes for player data
	r.Post("/player", ph.Create)
	r.Get("/player", ph.Get)
	r.Put("/player", ph.Update)
	r.Delete("/player", ph.Delete)

	// set routes for admin data
	r.Get("/admin", admin.Get)
	r.Post("/admin", admin.Update)

	// set routes for team data
	r.Get("/team", team.Get)
	r.Post("/team", team.Update)

	// set routes for challenge data
	r.Get("/challenge", challenge.Get)
	r.Post("/challenge", challenge.Update)

	// set routes for scoreboard data
	r.Get("/scoreboard", scoreboard.Get)
	r.Post("/scoreboard", scoreboard.Update)
}
