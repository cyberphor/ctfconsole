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
	// get api
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// wire health handlers to health endpoint
	v1.Get("/ruok", health.Get)

	// wire player handlers to player endpoint
	v1.Post("/player", player.Post)
	v1.Get("/player", player.Get)
	v1.Put("/player", player.Put)
	v1.Delete("/player", player.Delete)

	// wire admin handlers to admin endpoint
	v1.Post("/admin", admin.Post)
	v1.Get("/admin", admin.Get)
	v1.Put("/admin", admin.Put)
	v1.Delete("/admin", admin.Delete)

	// wire team handlers to team endpoint
	v1.Post("/team", team.Post)
	v1.Get("/team", team.Get)
	v1.Put("/team", team.Put)
	v1.Delete("/team", team.Delete)

	//  wire challenge handlers to challenge endpoint
	v1.Post("/challenge", challenge.Post)
	v1.Get("/challenge", challenge.Get)
	v1.Put("/challenge", challenge.Put)
	v1.Delete("/challenge", challenge.Delete)

	// wire scoreboard handlers to scoreboard endpoint
	v1.Post("/scoreboard", scoreboard.Post)
	v1.Get("/scoreboard", scoreboard.Get)
	v1.Put("/scoreboard", scoreboard.Put)
	v1.Delete("/scoreboard", scoreboard.Delete)

	err := app.Listen(":8081")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
