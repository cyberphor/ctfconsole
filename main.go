package main

import (
	"github.com/cyberphor/ctfconsole/controllers"
	"github.com/cyberphor/ctfconsole/models"
)

func main() {
	models.CreateTableForAdmins()
	models.CreateTableForPlayers()
	models.CreateAdmin("elliot", "bug")
	models.CreatePlayer("victor", "password", "Blue Team")
	controllers.Console()
}
