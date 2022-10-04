package service

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(p *dto.Account) (err error) {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(p.Secret), 8)

	p.Secret = string(hashed)

	err = repository.StoreAccountDB(p)
	return err

}
