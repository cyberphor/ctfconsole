package store

import (
	"database/sql"
)

type Driver string

const (
	Memory     Driver = "memory"
	SQLite3    Driver = "sqlite3"
	PostgreSQL Driver = "postgres"
)

type Store struct {
	Driver      Driver
	Name        string `default:"ctfconsole"`
	Address     string
	DB          *sql.DB
	Credentials string
}

func New() *Store {
	var s *Store
	var path string
	var err error

	s = &Store{
		Driver: SQLite3,
		Name:   "ctfconsole.db",
	}
	path = "storage/" + s.Name

	if s.Driver == Memory {
		// TODO: add memory as a storage option
		if err != nil {
			panic(err)
		}
	}

	if s.Driver == SQLite3 {
		s.DB, err = sql.Open(string(s.Driver), path)
		if err != nil {
			panic(err)
		}
	}
	return s
}
