package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

var (
	DB *sql.DB
)

func setupDB() error {
	var err error
	DB, err = sql.Open("postgres", "postgres://dummy:dummy@127.0.0.1:5432/notebook")
	if err != nil {
		return errors.Wrap(err, "failed to connect to the DB")
	}

	_, err = DB.Exec("DELETE FROM contacts;")
	if err != nil {
		return errors.Wrap(err, "failed to clear the contacts table")
	}

	return nil
}

func main() {
	if err := setupDB(); err != nil {
		log.Fatal(err)
	}
}
