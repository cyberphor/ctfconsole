package main

import (
	"github.com/cyberphor/ctfconsole/models"
	"github.com/cyberphor/ctfconsole/views"
)

func main() {
	models.CreateTableForUsers()
	models.CreateTableForAdmins()
	models.CreateAdmin("admin", "password", "admin")
	views.Console()
}
