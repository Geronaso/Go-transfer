package main

import (
	"go-transfer/internal/handler"
	"go-transfer/internal/repository"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {

	err := repository.StartDb()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e = handler.NewRouter(e)

	e.Logger.Fatal(e.Start(":8080"))

}
