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
