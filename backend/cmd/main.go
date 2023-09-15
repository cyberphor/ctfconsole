package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/cyberphor/ctfconsole/pkg/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetAddress(ip string, port int) string {
	return "http://" + ip + ":" + strconv.Itoa(port)
}

func Logger(logFilePath string) (*slog.Logger, error) {
	var file *os.File
	var err error
	var writer io.Writer
	var handler slog.Handler

	file, err = os.Create(logFilePath)
	writer = io.Writer(file)
	handler = slog.NewJSONHandler(writer, nil)
	return slog.New(handler), err
}

func main() {
	var uiIp string
	var uiPort int
	var apiIp string
	var apiPort int
	var apiLogPath string
	var err error
	var app *fiber.App
	var db *store.Store
	var logger *slog.Logger

	// get configuration arguments
	uiIp = *flag.String("ui-ip", "localhost", "UI IP Address")
	uiPort = *flag.Int("ui-port", 443, "UI Port")
	apiIp = *flag.String("api-ip", "localhost", "API IP Address")
	apiPort = *flag.Int("api-port", 8443, "API Port")
	apiLogPath = *flag.String("log-path", "/var/log/ctfconsole/ctfconsole.log", "Log file path")
	flag.Parse()

	// get logger
	logger, err = Logger(apiLogPath)
	if err != nil {
		panic(err)
	}
	logger.Info("Started logger")

	app = fiber.New()

	// allow inbound requests to backend from frontend
	app.Use(cors.New(cors.Config{AllowOrigins: GetAddress(uiIp, uiPort)}))
	db = store.New()
	router.Route(app, db)
	err = app.Listen(GetAddress(apiIp, apiPort))
	if err != nil {
		logger.Error(err.Error())
	}
}
