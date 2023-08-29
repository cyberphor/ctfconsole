package router

import (
	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/scoreboard"
	"github.com/cyberphor/ctfconsole/pkg/store"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, s *store.Store) {
	var r fiber.Router
	var ph *player.Handler

	r = app.Group("/api/v1")

	ph = &player.Handler{
		Store: s,
	}

	// player routes
	r.Post("/player", ph.Create)
	r.Get("/player", ph.Get)
	r.Put("/player", ph.Update)
	r.Delete("/player", ph.Delete)

	// admin routes
	r.Get("/api/v1/admin", admin.Get)
	r.Get("/api/v1/admin/:name", admin.Get)
	r.Post("/api/v1/admin", admin.Update)

	// team routes
	r.Get("/api/v1/team", team.Get)
	r.Get("/api/v1/team/:name", team.Get)
	r.Post("/api/v1/team", team.Update)

	// challenge routes
	r.Get("/api/v1/challenge", challenge.Get)
	r.Get("/api/v1/challenge/:name", challenge.Get)
	r.Post("/api/v1/challenge", challenge.Update)

	// scoreboard routes
	r.Get("/api/v1/scoreboard", scoreboard.Get)
	r.Get("/api/v1/scoreboard/{scoreboardId}", scoreboard.Get)
	r.Post("/api/v1/scoreboard", scoreboard.Update)
}
