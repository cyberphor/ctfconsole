package main

import (
	"github.com/cyberphor/ctfconsole/create"
	"github.com/cyberphor/ctfconsole/serve"
)

func main() {
	create.TableForAdmins()
	create.TableForPlayers()
	create.Admin("elliot", "bug")
	create.Player("victor", "password", "Blue Team")
	serve.Console()
}
