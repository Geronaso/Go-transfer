package tests

import (
	"go-transfer/internal/handler"
	"go-transfer/internal/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
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

func TestHandler_postacc(t *testing.T) {
	var testCases = []struct {
		name         string
		whenURL      string
		whenBody     string
		expectBody   string
		expectStatus int
	}{
		{
			name:         "ok",
			whenURL:      "/accounts",
			whenBody:     `{"name": "Paulo","cpf": "12345678902","secret": "123456"}`,
			expectStatus: http.StatusOK,
			expectBody:   "Account Created",
		},
		{
			name:         "ok",
			whenURL:      "/accounts",
			whenBody:     `{"name": "Ricardo","cpf": "12345678902","secret": "123456"}`,
			expectStatus: http.StatusBadRequest,
			expectBody:   "{\"message\":\"the cpf is already in use\"}\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repository.StartDb()
			if err != nil {
				log.Fatal(err)
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			e.POST("/accounts", handler.PostAcc)

			req := httptest.NewRequest(http.MethodPost, tc.whenURL, strings.NewReader(tc.whenBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectStatus, rec.Code)
			assert.Equal(t, tc.expectBody, rec.Body.String())
		})
	}
}

func TestHandler_getacc(t *testing.T) {
	var testCases = []struct {
		name         string
		whenURL      string
		whenBody     string
		expectBody   string
		expectStatus int
	}{
		{
			name:         "ok",
			whenURL:      "/accounts/Paulo/balance",
			expectStatus: http.StatusOK,
			expectBody:   "1000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repository.StartDb()
			if err != nil {
				log.Fatal(err)
			}

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New()}

			e.GET("/accounts/:id/balance", handler.GetAccId)

			req := httptest.NewRequest(http.MethodGet, tc.whenURL, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectStatus, rec.Code)
			assert.Equal(t, tc.expectBody, rec.Body.String())
		})
	}
}
