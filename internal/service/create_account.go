package service

import (
	"go-transfer/internal/datastruct"
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(account *dto.Account) (err error) {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(account.Secret), 8)

	accountDB := new(datastruct.Account)

	accountDB.Secret = string(hashed)
	accountDB.Name = account.Name
	accountDB.Cpf = account.Cpf
	accountDB.Created_at = account.Created_at
	accountDB.Balance = account.Balance

	err = repository.StoreAccountDB(accountDB)
	return err

}
