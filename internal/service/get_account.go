package service

import (
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"
)

func GetAccounts() ([]dto.AccountGet, error) {

	accounts, err := repository.GetAccountsDB()

	accountsDTO := make([]dto.AccountGet, len(accounts))

	for i, v := range accounts {
		accountsDTO[i].Name = v.Name
		accountsDTO[i].Cpf = v.Cpf
		accountsDTO[i].Balance = v.Balance
		accountsDTO[i].Created_at = v.Created_at
	}

	return accountsDTO, err

}

func GetBalance(user string) (string, error) {

	balance, err := repository.GetBalanceDB(user)
	return balance, err

}
