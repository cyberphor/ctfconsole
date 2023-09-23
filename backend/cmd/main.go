package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
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
	return c.APIIP + ":" + strconv.Itoa(c.APIPort)
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

type Health struct {
	Status string `default:"imok"`
}

func (h Health) Get(c *fiber.Ctx) error {
	var message map[string]string

	message = make(map[string]string)
	message["data"] = h.Status
	fmt.Println(message)
	return c.Status(200).JSON(message)
}

func getEnvStr(key string) (string, error) {
	var valueStr string
	var defined bool

	valueStr, defined = os.LookupEnv(key)
	if !defined {
		return valueStr, fmt.Errorf("%s is not defined", key)
	}

	return valueStr, nil
}

func getEnvInt(key string) (int, error) {
	var valueStr string
	var valueInt int
	var defined bool
	var err error

	valueStr, defined = os.LookupEnv(key)
	if !defined {
		err = errors.New("")
		return 0, fmt.Errorf("%s is not defined", key)
	}

	valueInt, err = strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("%s is not an integer", key)
	}

	return valueInt, nil
}

func GetConfig() (Config, error) {
	var config Config
	var err error

	// get ui parameters
	config.UIProtocol, err = getEnvStr("CTFCONSOLE_UI_PROTOCOL")
	config.UIIP, err = getEnvStr("CTFCONSOLE_UI_IP_ADDRESS")
	config.UIPort, err = getEnvInt("CTFCONSOLE_UI_PORT")

	// get api parameters
	config.APILogPath, err = getEnvStr("CTFCONSOLE_API_LOG_PATH")
	config.APIProtocol, err = getEnvStr("CTFCONSOLE_API_PROTOCOL")
	config.APIIP, err = getEnvStr("CTFCONSOLE_API_IP_ADDRESS")
	config.APIPort, err = getEnvInt("CTFCONSOLE_API_PORT")

	// get db parameters
	config.DBName, err = getEnvStr("CTFCONSOLE_DB_NAME")
	config.DBUsername, err = getEnvStr("CTFCONSOLE_DB_USER")
	config.DBPassword, err = getEnvStr("CTFCONSOLE_DB_PASSWORD")
	config.DBProtocol, err = getEnvStr("CTFCONSOLE_DB_PROTOCOL")
	config.DBIP, err = getEnvStr("CTFCONSOLE_DB_IP_ADDRESS")
	config.DBPort, err = getEnvInt("CTFCONSOLE_DB_PORT")

	return config, err
}

func main() {
	var config Config
	var err error
	var logger *slog.Logger
	var app *fiber.App

	// get config
	config, err = GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get api logger
	logger, err = Logger(config.APILogPath)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	logger.Info("Started logger")

	// configure api to accept inbound requests from ui
	app = fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	// get db connection
	//db, err = sql.Open("postgres", config.GetDBEndpoint())
	//if err != nil {
	//	fmt.Println(err.Error())
	//	logger.Error(err.Error())
	//}

	// wire api to db

	var health *Health
	//var ph *player.Handler

	// get health handler
	health = &Health{}

	// get player handler
	//ph = &player.Handler{
	//	DB: db,
	//}

	// set routes for api health data
	app.Get("/api/v1/ruok", health.Get)

	// // set routes for player data
	// app.Post("/api/v1/player", ph.Create)
	// app.Get("/api/v1/player", ph.Get)
	// app.Put("/api/v1/player", ph.Update)
	// app.Delete("/api/v1/player", ph.Delete)
	//
	// // set routes for admin data
	// app.Get("/api/v1/admin", admin.Get)
	// app.Post("/api/v1/admin", admin.Update)
	//
	// // set routes for team data
	// app.Get("/team", team.Get)
	// app.Post("/team", team.Update)
	//
	// // set routes for challenge data
	// app.Get("/challenge", challenge.Get)
	// app.Post("/challenge", challenge.Update)
	//
	// // set routes for scoreboard data
	// app.Get("/scoreboard", scoreboard.Get)
	// app.Post("/scoreboard", scoreboard.Update)

	err = app.Listen("127.0.0.1:8443")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		// logger.Error(err.Error())
	}
}
