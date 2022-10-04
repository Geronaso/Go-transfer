package service

import (
	"go-transfer/internal/datastruct"
	"go-transfer/internal/repository"
)

func GetAccounts() ([]datastruct.Account, error) {

	accounts, err := repository.GetAccountsDB()
	return accounts, err

}

func GetBalance(user string) (string, error) {

	balance, err := repository.GetBalanceDB(user)
	return balance, err

}
