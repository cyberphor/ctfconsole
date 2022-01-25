package main

import (
	"github.com/cyberphor/ctfconsole/create"
	"github.com/cyberphor/ctfconsole/list"
	"github.com/cyberphor/ctfconsole/serve"
)

func main() {
	create.TableForAdmins()
	create.TableForPlayers()
	create.CreateAdmin("elliot", "bug")
	create.CreatePlayer("victor", "password")
	list.ListPlayers()
	serve.Console()
}
