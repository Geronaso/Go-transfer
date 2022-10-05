package handler

import (
	"go-transfer/internal/datastruct"
	"go-transfer/internal/dto"
	"go-transfer/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostLog(c echo.Context) (err error) {

	login_info := new(dto.Login)

	// Validate data received
	if err := c.Bind(login_info); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(login_info); err != nil {
		return err
	}

	token := new(datastruct.Token)
	// Store on Database
	if token.Token, err = service.ValidateUser(login_info); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSONPretty(http.StatusOK, token, "")
}
