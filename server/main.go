package main

import (
	"github.com/cyberphor/ctfconsole/handlers"
	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Team  string `json:"team"`
}

func main() {
	var app *fiber.App = fiber.New()
	app.Get("/", handlers.Index)
	app.Get("/api/player/{playerId}", handlers.GetPlayer)
	app.Get("/api/players", handlers.GetPlayers)
	app.Get("/api/admin/{adminId}", handlers.GetAdmin)
	app.Get("/api/admins", handlers.GetAdmins)
	app.Get("/api/team/{teamId}", handlers.GetTeam)
	app.Get("/api/teams", handlers.GetTeams)
	app.Get("/api/challenge/{challengeId}", handlers.GetChallenge)
	app.Get("/api/challenges", handlers.GetChallenges)
	app.Get("/api/scoreboard/{scoreboardId}", handlers.GetScoreboard)
	app.Get("/api/scoreboards", handlers.GetScoreboards)

	app.Post("/api/player", handlers.UpdatePlayer)
	app.Post("/api/admin", handlers.UpdateAdmin)
	app.Post("/api/team", handlers.UpdateAdmin)
	app.Post("/api/challenge", handlers.UpdateChallenge)
	app.Post("/api/scoreboard", handlers.UpdateScoreboard)
	app.Listen(":9001")
}
