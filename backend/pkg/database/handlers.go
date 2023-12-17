package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	// get db host
	host, defined := os.LookupEnv("CTFCONSOLE_DB_HOST")
	if !defined {
		return nil, fmt.Errorf("CTFCONSOLE_DB_HOST is not defined")
	}

	// get db port
	port, defined := os.LookupEnv("CTFCONSOLE_DB_PORT")
	if !defined {
		return nil, fmt.Errorf("CTFCONSOLE_DB_PORT is not defined")
	}

	// get db name
	name, defined := os.LookupEnv("CTFCONSOLE_DB_NAME")
	if !defined {
		return nil, fmt.Errorf("CTFCONSOLE_DB_NAME is not defined")
	}

	// get db user
	user, defined := os.LookupEnv("CTFCONSOLE_DB_USER")
	if !defined {
		return nil, fmt.Errorf("CTFCONSOLE_DB_USER is not defined")
	}

	// get db password
	password, defined := os.LookupEnv("CTFCONSOLE_DB_PASSWORD")
	if !defined {
		return nil, fmt.Errorf("CTFCONSOLE_DB_PASSWORD is not defined")
	}

	// get db connection string
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)

	// connect to db
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// check if db connection works
	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
