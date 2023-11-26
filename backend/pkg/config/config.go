package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port    int
	LogPath string
}

func Get() (Config, error) {
	var config Config
	var defined bool
	var err error
	var logpath string
	var port string

	// get API port
	port, defined = os.LookupEnv("CTFCONSOLE_API_PORT")
	if !defined {
		return config, fmt.Errorf("'CTFCONSOLE_API_PORT' is not defined")
	}

	// convert API port value to an int
	config.Port, err = strconv.Atoi(port)
	if err != nil {
		return config, err
	}

	// get API logpath value
	logpath, defined = os.LookupEnv("CTFCONSOLE_LOG_PATH")
	if !defined {
		return config, fmt.Errorf("'CTFCONSOLE_LOG_PATH' is not defined")
	}
	config.LogPath = logpath

	// return config
	return config, nil
}

func DatabaseConnection() (*sql.DB, error) {
	username, defined := os.LookupEnv("CTFCONSOLE_DB_USERNAME")
	if !defined {
		return nil, fmt.Errorf("'CTFCONSOLE_DB_USERNAME' is not defined")
	}

	password, defined := os.LookupEnv("CTFCONSOLE_DB_PASSWORD")
	if !defined {
		return nil, fmt.Errorf("'CTFCONSOLE_DB_PASSWORD' is not defined")
	}

	address, defined := os.LookupEnv("CTFCONSOLE_DB_ADDRESS")
	if !defined {
		return nil, fmt.Errorf("'CTFCONSOLE_DB_ADDRESS' is not defined")
	}

	name, defined := os.LookupEnv("CTFCONSOLE_DB_NAME")
	if !defined {
		return nil, fmt.Errorf("'CTFCONSOLE_DB_NAME' is not defined")
	}

	connection := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		address,
		name,
	)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}
