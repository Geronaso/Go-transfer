package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var db *sql.DB
var err error

func CheckErr(err error) error {

	if err != nil {
		string_error := err.Error()
		fmt.Println(string_error)

		if strings.Contains(string_error, "Rows are closed") {
			return errors.New("not found on database")
		}
		if strings.Contains(string_error, "UNIQUE constraint failed: account.cpf") {
			return errors.New("the cpf is already in use")
		}

	}
	return err
}

func StartDb() error {
	// Connect to database
	db, err = sql.Open("sqlite3", "./database.db")
	CheckErr(err)

	return db.Ping()
}
