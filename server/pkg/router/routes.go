package router

import (
	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/scoreboard"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
)

func Set(app *fiber.App) {
	// default route
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Welcome") })

	// player routes
	app.Post("/api/v1/player/:name/:password", player.Create)
	app.Get("/api/v1/player", player.Get)
	app.Get("/api/v1/player/:name", player.Get)
	app.Put("/api/v1/player/:name", player.Get)
	app.Delete("/api/v1/player/:username", player.Get)

	// admin routes
	app.Get("/api/v1/admin", admin.Get)
	app.Get("/api/v1/admin/:name", admin.Get)
	app.Post("/api/v1/admin", admin.Update)

	// team routes
	app.Get("/api/v1/team", team.Get)
	app.Get("/api/v1/team/:name", team.Get)
	app.Post("/api/v1/team", team.Update)

	// challenge routes
	app.Get("/api/v1/challenge", challenge.Get)
	app.Get("/api/v1/challenge/:name", challenge.Get)
	app.Post("/api/v1/challenge", challenge.Update)

	// scoreboard routes
	app.Get("/api/v1/scoreboard", scoreboard.Get)
	app.Get("/api/v1/scoreboard/{scoreboardId}", scoreboard.Get)
	app.Post("/api/v1/scoreboard", scoreboard.Update)
}
