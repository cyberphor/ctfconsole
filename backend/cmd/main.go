package main

import (
	"database/sql"
	"flag"
	"io"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetURL(protocol string, ip string, port int) string {
	return protocol + "://" + ip + ":" + strconv.Itoa(port)
}

func GetDataSourceName(protocol string, username string, password string, ip string, port int, database string) string {
	return protocol + "://" + username + ":" + password + "@" + ip + "/" + database + "?sslmode=disable"
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
	var dbIp string
	var dbPort int
	var dbUsername string
	var dbPassword string
	var dbName string
	var err error
	var logger *slog.Logger
	var app *fiber.App
	var db *sql.DB
	var dataSourceName string

	// get configuration arguments
	uiIp = *flag.String("ui-ip", "localhost", "UI IP Address")
	uiPort = *flag.Int("ui-port", 443, "UI Port")
	apiIp = *flag.String("api-ip", "localhost", "API IP Address")
	apiPort = *flag.Int("api-port", 8443, "API Port")
	apiLogPath = *flag.String("log-path", "/var/log/ctfconsole/ctfconsole.log", "Log file path")
	dbIp = *flag.String("db-ip", "localhost", "DB IP Address")
	dbPort = *flag.Int("db-port", 5432, "DB Port")
	dbUsername = *flag.String("db-user", "postgres", "DB service account username")
	dbPassword = *flag.String("db-password", "postgres", "DB service account password")
	dbName = *flag.String("db-name", "ctfconsole", "DB name")
	flag.Parse()

	// get logger
	logger, err = Logger(apiLogPath)
	if err != nil {
		panic(err)
	}
	logger.Info("Started logger")

	// allow inbound requests to backend from frontend
	app = fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: GetURL("http", uiIp, uiPort)}))

	//
	dataSourceName = GetDataSourceName("postgres", dbUsername, dbPassword, dbIp, dbPort, dbName)
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	router.Route(app, db)
	err = app.Listen(GetURL("http", apiIp, apiPort))
	if err != nil {
		logger.Error(err.Error())
	}
}
