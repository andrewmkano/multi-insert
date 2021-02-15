package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Contact struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Contacts []Contact

func (cs Contacts) BulkInsert(db *sql.DB) error {
	var (
		placeholders []string
		vals         []interface{}
	)

	for index, contact := range cs {
		placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d)",
			index*3+1,
			index*3+2,
			index*3+3,
		))

		vals = append(vals, contact.FirstName, contact.LastName, contact.Email)
	}

	txn, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "could not start a new transaction")
	}

	insertStatement := fmt.Sprintf("INSERT INTO contacts(first_name,last_name,email) VALUES %s", strings.Join(placeholders, ","))
	_, err = txn.Exec(insertStatement, vals...)
	if err != nil {
		txn.Rollback()
		return errors.Wrap(err, "failed to insert multiple records at once")
	}

	if err := txn.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
