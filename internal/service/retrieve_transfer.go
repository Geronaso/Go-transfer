package service

import (
	"go-transfer/internal/datastruct"
	"go-transfer/internal/repository"
)

func RetrieveTransfer(cpf string) ([]datastruct.Transfer, error) {

	account, err := repository.RetrieveAccount(cpf)
	if err != nil {
		return nil, err
	}

	transfers, err := repository.RetrieveTransferDB(account)

	return transfers, err

}
