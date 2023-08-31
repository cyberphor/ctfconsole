package main

import (
	"flag"
	"io"
	"log/slog"
	"os"

	"github.com/cyberphor/ctfconsole/pkg/router"
	"github.com/cyberphor/ctfconsole/pkg/store"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type StoreCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	UiIp             string           `yaml:"uiIp"`
	UiPort           string           `yaml:"uiPort"`
	StoreDriver      string           `yaml:"storeDriver"`
	StoreName        string           `yaml:"storeName"`
	StoreIP          string           `yaml:"storeIp"`
	StorePort        string           `yaml:"storePort"`
	StoreCredentials StoreCredentials `yaml:"storeCredentials"`
	LogFilePath      string           `yaml:"logFilePath"`
}

func (c *Config) GetUiAddress() string {
	return c.UiIp + ":" + c.UiPort
}

func (c *Config) GetStoreAddress() string {
	return c.StoreIP + ":" + c.StorePort
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
	var configFilePath string
	var configFile []byte
	var err error
	var config Config
	var app *fiber.App
	var db *store.Store
	var logger *slog.Logger

	// get config file path
	configFilePath = *flag.String("c", "config.yaml", "Path to ctfconsole config")
	flag.Parse()

	// get config file
	configFile, err = os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	// get config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	// get logger
	logger, err = Logger(config.LogFilePath)
	if err != nil {
		panic(err)
	}
	logger.Info("Started logger")

	app = fiber.New()
	db = store.New()
	router.Route(app, db)
	err = app.Listen(config.GetUiAddress())
	if err != nil {
		logger.Error(err.Error())
	}
}
