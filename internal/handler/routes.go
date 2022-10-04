package handler

import "github.com/labstack/echo/v4"

func NewRouter(e *echo.Echo) *echo.Echo {

	e.GET("/accounts", GetAcc)
	e.GET("/accounts/:id/balance", GetAccId)
	e.POST("/accounts", PostAcc)
	// e.POST("/login", GetHome)
	// e.GET("/transfers", PostUser)
	// e.POST("/transfers", PostUser)

	return e
}
