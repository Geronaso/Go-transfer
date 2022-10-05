package handler

import (
	"go-transfer/internal/dto"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetTra(c echo.Context) (err error) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Cpf

	return c.String(http.StatusOK, "Welcome "+name+"!")

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

	return c.String(http.StatusOK, "Transfer Completed")

}
