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

type Health struct {
	Status string
}

func (h Health) Get(c *fiber.Ctx) error {
	var message map[string]string
	message = make(map[string]string)
	h.Status = "imok"
	message["data"] = h.Status
	return c.Status(200).JSON(message)
}

func Route(app *fiber.App, s *store.Store) {
	var r fiber.Router
	var health *Health
	var ph *player.Handler

	r = app.Group("/api/v1")
	health = &Health{}
	ph = &player.Handler{
		Store: s,
	}

	// health routes
	app.Get("/ruok", health.Get)

	// player routes
	r.Post("/player", ph.Create)
	r.Get("/player", ph.Get)
	r.Put("/player", ph.Update)
	r.Delete("/player", ph.Delete)

	// admin routes
	r.Get("/admin", admin.Get)
	r.Post("/admin", admin.Update)

	// team routes
	r.Get("/team", team.Get)
	r.Post("/team", team.Update)

	// challenge routes
	r.Get("/challenge", challenge.Get)
	r.Post("/challenge", challenge.Update)

	// scoreboard routes
	r.Get("/scoreboard", scoreboard.Get)
	r.Post("/scoreboard", scoreboard.Update)
}
