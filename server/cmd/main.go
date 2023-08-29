package main

import (
	"flag"
	"io/ioutil"

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
}

func (c *Config) GetUiAddress() string {
	return c.UiIp + ":" + c.UiPort
}

func (c *Config) GetStoreAddress() string {
	return c.StoreIP + ":" + c.StorePort
}

func main() {
	var configFilePath string
	var configFile []byte
	var err error
	var config Config
	var app *fiber.App
	var db *store.Store

	// get config file path
	configFilePath = *flag.String("c", "config.yaml", "Path to ctfconsole config")
	flag.Parse()

	// get config file
	configFile, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	// get config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	app = fiber.New()
	db = store.New()
	router.Route(app, db)
	err = app.Listen(config.GetUiAddress())
	if err != nil {
		panic(err)
	}
}
