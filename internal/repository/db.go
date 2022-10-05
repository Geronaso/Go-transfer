package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func CheckErr(err error) error {

	if err != nil {
		string_error := err.Error()
		fmt.Println(string_error)

		if strings.Contains(string_error, "Rows are closed") {
			return errors.New("not found on database")
		}

	}
	return err
}

func StartDb() (*sql.DB, error) {
	// Connect to database
	db, err := sql.Open("sqlite3", "./database.db")
	CheckErr(err)

	return db, err
}
