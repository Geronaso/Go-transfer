package repository

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckErr(err error) error {

	if err != nil {
		string_error := err.Error()

		if strings.Contains(string_error, "sql: no rows in result set") {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
