package main

import (
	"github.com/cyberphor/ctfconsole/controllers"
	"github.com/cyberphor/ctfconsole/models"
)

func main() {
	models.CreateTableForUsers()
	models.CreateTableForAdmins()
	models.CreateAdmin("admin", "password", "admin")
	models.CreateTableForScoreboard()
	controllers.Console()
}
