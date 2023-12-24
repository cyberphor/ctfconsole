package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/cyberphor/ctfconsole/pkg/admin"
	"github.com/cyberphor/ctfconsole/pkg/campaign"
	"github.com/cyberphor/ctfconsole/pkg/challenge"
	"github.com/cyberphor/ctfconsole/pkg/health"
	"github.com/cyberphor/ctfconsole/pkg/player"
	"github.com/cyberphor/ctfconsole/pkg/team"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// health, admin, campaign, challenge, team, player

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// get app
	app := fiber.New()

	// get app logpath
	logpath, defined := os.LookupEnv("CTFCONSOLE_API_LOG_PATH")
	if !defined {
		logger.Error("CTFCONSOLE_API_LOG_PATH is not defined")
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Logpath: %s", logpath))

	// get app router
	router := app.Group("/api/v1")

	// wire health endpoint to health handlers
	router.Get("/ruok", health.Get)

	// wire admin endpoint to admin handlers
	router.Post("/admin", admin.Post)
	router.Get("/admin", admin.Get)
	router.Put("/admin", admin.Put)
	router.Delete("/admin", admin.Delete)

	// wire campaign endpoint to campaign handlers
	router.Post("/campaign", campaign.Post)
	router.Get("/campaign", campaign.Get)
	router.Put("/campaign", campaign.Put)
	router.Delete("/campaign", campaign.Delete)

	// wire challenge endpoint to challenge handlers
	router.Post("/challenge", challenge.Post)
	router.Get("/challenge", challenge.Get)
	router.Put("/challenge", challenge.Put)
	router.Delete("/challenge", challenge.Delete)

	// wire team endpoint to team handlers
	router.Post("/team", team.Post)
	router.Get("/team", team.Get)
	router.Put("/team", team.Put)
	router.Delete("/team", team.Delete)

	// wire player endpoint to player handlers
	router.Post("/player", player.Post)
	router.Get("/player", player.Get)
	router.Put("/player", player.Put)
	router.Delete("/player", player.Delete)

	// get app address
	port, defined := os.LookupEnv("CTFCONSOLE_API_PORT")
	if !defined {
		logger.Error("CTFCONSOLE_API_PORT is not defined")
		os.Exit(1)
	}
	address := fmt.Sprintf(":%s", port)

	// start app
	err := app.Listen(address)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
