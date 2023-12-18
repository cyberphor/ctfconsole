package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	// get db host
	host, defined := os.LookupEnv("POSTGRES_HOST")
	if !defined {
		return nil, fmt.Errorf("POSTGRES_HOST is not defined")
	}

	// get db port
	port, defined := os.LookupEnv("POSTGRES_PORT")
	if !defined {
		return nil, fmt.Errorf("POSTGRES_PORT is not defined")
	}

	// get db name
	name, defined := os.LookupEnv("POSTGRES_DB")
	if !defined {
		return nil, fmt.Errorf("POSTGRES_DB is not defined")
	}

	// get db user
	user, defined := os.LookupEnv("POSTGRES_USER")
	if !defined {
		return nil, fmt.Errorf("POSTGRES_USER is not defined")
	}

	// get db password
	password, defined := os.LookupEnv("POSTGRES_PASSWORD")
	if !defined {
		return nil, fmt.Errorf("POSTGRES_PASSWORD is not defined")
	}

	// get db connection string
	dataSourceName := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
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
