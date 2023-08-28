package store

import (
	"database/sql"
)

type Type string

const (
	Memory     Type = "memory"
	SQLite3    Type = "sqlite3"
	PostgreSQL Type = "postgres"
)

type Config struct {
	Type        Type
	Name        string `default:"ctfconsole"`
	Address     string
	Credentials string
}

type Store struct {
	Type        Type
	Name        string `default:"ctfconsole"`
	Address     string
	Credentials string
	DB          *sql.DB
}

func New(config ...Config) *Store {
	var s *Store
	var path string
	var err error

	s = &Store{
		Type:        config.Type,
		Name:        config.Name,
		Address:     config.Address,
		Credentials: config.Credentials,
	}

	path = "storage/" + s.Name
	if s.Type == SQLite3 {
		s.DB, err = sql.Open(string(s.Type), path)
		if err != nil {
			panic(err)
		}
	} else if s.Type == PostgreSQL {
		s.DB, err = sql.Open(string(s.Type), path)
		if err != nil {
			panic(err)
		}
	}
	return s
}

/*
  memory
  sqlite3
	-

	postgres
*/
