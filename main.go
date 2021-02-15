package main

import (
	"database/sql"
	"fmt"
	"log"
	"multi-insert/generator"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	DB              *sql.DB
	recordsToInsert = int(100)
)

func setupDB() error {
	var err error

	DB, err = sql.Open("postgres", "postgres://dummy:dummy@127.0.0.1:5432/notebook?sslmode=disable")
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

	dummyContacts := generator.GenerateDummyContacts(recordsToInsert)

	// Just to get an idea of how long its taking us to insert that many records (roughly speaking)
	n := time.Now()

	if err := dummyContacts.BulkInsert(DB); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Inserting %d took %s", recordsToInsert, time.Since(n)))
}
