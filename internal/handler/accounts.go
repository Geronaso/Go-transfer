package handler

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PostAcc(c echo.Context) (err error) {

	account := new(dto.Account)

	current_time := time.Now().Local().Format(time.RFC3339)
	account.Created_at = current_time
	account.Balance = 0

	// Validate data received
	if err = c.Bind(account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(account); err != nil {
		return err
	}

	// Store on Database
	if err = service.CreateAccount(account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Account Created")
}

func GetAcc(c echo.Context) (err error) {

	accounts, err := service.GetAccounts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSONPretty(http.StatusOK, accounts, "")
}

func GetAccId(c echo.Context) (err error) {

	user := c.Param("id")

	balance, err := service.GetBalance(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, balance)
}
