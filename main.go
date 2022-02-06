package main

import (
	"github.com/cyberphor/ctfconsole/controllers"
	"github.com/cyberphor/ctfconsole/models"
)

func main() {
	models.CreateTableForUsers()
	models.CreateTableForScoreboard()
	controllers.Console()
}
