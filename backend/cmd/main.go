package main

import (
	"database/sql"
	"flag"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Config struct {
	UIProtocol  string
	UIIP        string
	UIPort      int
	APIProtocol string
	APIIP       string
	APIPort     int
	APILogPath  string
	DBProtocol  string
	DBIP        string
	DBPort      int
	DBUsername  string
	DBPassword  string
	DBName      string
}

func (c Config) GetUIEndpoint() string {
	return c.UIProtocol + "://" + c.UIIP + ":" + strconv.Itoa(c.UIPort)
}

func (c Config) GetAPIEndpoint() string {
	return c.APIProtocol + "://" + c.APIIP + ":" + strconv.Itoa(c.APIPort)
}

func (c Config) GetDBEndpoint() string {
	return c.DBProtocol + "://" + c.DBUsername + ":" + c.DBPassword + "@" + c.DBIP + "/" + c.DBName + "?sslmode=disable"
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
	var config Config
	var err error
	var logger *slog.Logger
	var app *fiber.App
	var db *sql.DB

	// get ui parameters
	flag.StringVar(&config.UIProtocol, "ui-proto", "http", "UI protocol")
	flag.StringVar(&config.UIIP, "ui-ip", "localhost", "UI IP address")
	flag.IntVar(&config.UIPort, "ui-port", 443, "UI Port")

	// get api parameters
	flag.StringVar(&config.APIProtocol, "api-proto", "http", "API protocol")
	flag.StringVar(&config.APIIP, "api-ip", "localhost", "API IP address")
	flag.IntVar(&config.APIPort, "api-port", 8443, "API Port")
	flag.StringVar(&config.APILogPath, "log-path", "/var/log/ctfconsole/ctfconsole.log", "Log file path")

	// get db parameters
	flag.StringVar(&config.DBProtocol, "db-proto", "localhost", "DB protocol")
	flag.StringVar(&config.DBIP, "db-ip", "localhost", "DB IP address")
	flag.IntVar(&config.DBPort, "db-port", 5432, "DB Port")
	flag.StringVar(&config.DBUsername, "db-user", "postgres", "DB service account username")
	flag.StringVar(&config.DBPassword, "db-password", "postgres", "DB service account password")
	flag.StringVar(&config.DBName, "db-name", "ctfconsole", "DB name")
	flag.Parse()

	// get api logger
	logger, err = Logger(config.APILogPath)
	if err != nil {
		panic(err)
	}
	logger.Info("Started logger")

	// configure api to accept inbound requests from ui
	app = fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: config.GetUIEndpoint()}))

	// get db connection
	db, err = sql.Open("postgres", config.GetDBEndpoint())
	if err != nil {
		logger.Error(err.Error())
	}

	// wire api to db
	router.Route(app, db)
	err = app.Listen(config.GetAPIEndpoint())
	if err != nil {
		logger.Error(err.Error())
	}
}
