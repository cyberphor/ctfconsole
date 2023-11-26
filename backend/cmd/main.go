package main

import (
	"fmt"

	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/config"
	"github.com/cyberphor/ctfconsole/pkg/health"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/scoreboard"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	// declare variables
	var c config.Config
	var err error

	// get ctfconsole config
	c, err = config.Get()
	if err != nil {
		panic(err)
	}

	// get api
	app := fiber.New()
	router := app.Group("/api/v1")

	// wire health handlers to health endpoint
	router.Get("/ruok", health.Get)

	// wire player handlers to player endpoint
	router.Post("/player", player.Post)
	router.Get("/player", player.Get)
	router.Put("/player", player.Put)
	router.Delete("/player", player.Delete)

	// wire admin handlers to admin endpoint
	router.Post("/admin", admin.Post)
	router.Get("/admin", admin.Get)
	router.Put("/admin", admin.Put)
	router.Delete("/admin", admin.Delete)

	// wire team handlers to team endpoint
	router.Post("/team", team.Post)
	router.Get("/team", team.Get)
	router.Put("/team", team.Put)
	router.Delete("/team", team.Delete)

	//  wire challenge handlers to challenge endpoint
	router.Post("/challenge", challenge.Post)
	router.Get("/challenge", challenge.Get)
	router.Put("/challenge", challenge.Put)
	router.Delete("/challenge", challenge.Delete)

	// wire scoreboard handlers to scoreboard endpoint
	router.Post("/scoreboard", scoreboard.Post)
	router.Get("/scoreboard", scoreboard.Get)
	router.Put("/scoreboard", scoreboard.Put)
	router.Delete("/scoreboard", scoreboard.Delete)

	err = app.Listen(fmt.Sprintf(":%d", c.Port))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
