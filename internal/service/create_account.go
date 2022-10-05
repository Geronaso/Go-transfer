package service

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(account *dto.Account) (err error) {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(account.Secret), 8)

	account.Secret = string(hashed)

	err = repository.StoreAccountDB(account)
	return err

}
