package handler

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetTra(c echo.Context) (err error) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	cpf := claims.Cpf

	transfers, err := service.RetrieveTransfer(cpf)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSONPretty(http.StatusOK, transfers, "")

}

func PostTra(c echo.Context) (err error) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	cpf := claims.Cpf
	current_time := time.Now().Local().Format(time.RFC3339)

	transfer := new(dto.Transfer)
	transfer.Account_origin = cpf
	transfer.Created_at = current_time

	// Validate data received
	if err = c.Bind(transfer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(transfer); err != nil {
		return err
	}

	if transfer.Account_destination == transfer.Account_origin {
		return echo.NewHTTPError(http.StatusBadRequest, "You can not make a transfer to yourself")
	}

	if err = service.ProcessTransfer(transfer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Transfer Completed")

}