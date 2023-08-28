package main

import (
	"flag"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/cyberphor/ctfconsole/pkg/store"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var ip string
	var port string
	var address string
	var app *fiber.App
	var db *store.Store

	ip = *flag.String("ip", "", "ctfconsole IP address")
	port = *flag.String("p", "80", "ctfconsole UI port")
	address = ip + ":" + port
	// store type
	// store name
	// store address
	// store credentials
	flag.Parse()

	app = fiber.New()
	db = store.New()
	router.Route(app, db)
	app.Listen(address)
}
