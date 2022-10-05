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
	hashed_pass, err := repository.RetrieveHashPass(user.Cpf)
	if err != nil {
		return hashed_pass, err
	}
	// Compare Hashed pass with Secret
	if err = bcrypt.CompareHashAndPassword([]byte(hashed_pass), []byte(user.Secret)); err != nil {
		return hashed_pass, err
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.Cpf,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return t, err
	}

	// bytes_token := generate_token()

	// // TO-DO STORE TOKEN, CREATE DATABASE TABLE THINKING ON THE TRANSFER DATABASE TABLE
	// // Hash the generated token and store on DB
	// hashed_token, _ := bcrypt.GenerateFromPassword([]byte(bytes_token), 8)

	// err = repository.StoreTokenDB(string(hashed_token), user.Cpf)

	return t, err
}

// // Generate a random token
// func generate_token() string {
// 	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
// 	str := make([]rune, 32)
// 	for i := range str {
// 		str[i] = chars[rand.Intn(len(chars))]
// 	}
// 	return string(str)
// }

type jwtCustomClaims struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}
