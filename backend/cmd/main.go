package main

import (
	"fmt"

	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/health"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/scoreboard"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	// get handlers
	health := health.Handler{}
	player := player.Handler{}

	// get api
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// wire health handlers to health endpoint
	v1.Get("/ruok", health.Get)

	// wire player handlers to player endpoint
	v1.Post("/player", player.Create)

	// wire admin handler to admin endpoint
	v1.Get("/admin", admin.Get)
	v1.Post("/admin", admin.Update)

	// wire team handler to team endpoint
	v1.Get("/team", team.Get)
	v1.Post("/team", team.Update)

	//  wire challenge handler to challenge endpoint
	v1.Get("/challenge", challenge.Get)
	v1.Post("/challenge", challenge.Update)

	// wire scoreboard handler to scoreboard endpoint
	v1.Get("/scoreboard", scoreboard.Get)
	v1.Post("/scoreboard", scoreboard.Update)

	err := app.Listen(":8081")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
