package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type client struct {
	DB *sql.DB
}

func NewClient(dir string) (client, error) {
	db, err := sql.Open("sqlite3", dir)
	if err != nil {
		return client{}, err
	}

	return client{db}, nil
}

func (c client) Close() error {
	return c.DB.Close()
}
