package service

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(user *dto.Login) (string, error) {

	// Retrieve hashed pass from database
	hashed_pass, err := repository.RetrieveHashPassDB(user.Cpf)
	if err != nil {
		return hashed_pass, err
	}
	// Compare Hashed pass with Secret
	if err = bcrypt.CompareHashAndPassword([]byte(hashed_pass), []byte(user.Secret)); err != nil {
		return hashed_pass, err
	}

	// Set custom claims for JWT expires after 24 hours

	type jwtCustomClaims struct {
		Cpf string `json:"cpf"`
		jwt.StandardClaims
	}

	claims := &jwtCustomClaims{
		user.Cpf,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	return t, err
}
