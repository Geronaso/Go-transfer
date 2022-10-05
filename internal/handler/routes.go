package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

func NewRouter(e *echo.Echo) *echo.Echo {

	e.GET("/accounts", GetAcc)
	e.GET("/accounts/:id/balance", GetAccId)
	e.POST("/accounts", PostAcc)
	e.POST("/login", PostLog)
	// e.GET("/transfers", GetTra)
	// e.POST("/transfers", PostTransf)

	// Restricted group
	r := e.Group("/transfers")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", GetTra)

	return e
}
