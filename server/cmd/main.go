package main

import (
	"flag"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/cyberphor/ctfconsole/pkg/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var port string
	var app *fiber.App

	port = ":" + *flag.String("p", "80", "ctfconsole UI port")
	flag.Parse()

	storage.CreateSQLiteDatabase("ctfconsole.db")
	app = fiber.New()
	router.Set(app)
	app.Listen(port)
}

// go build -o ctfconsole .
