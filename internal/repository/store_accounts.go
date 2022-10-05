package repository

import (
	"go-transfer/internal/dto"
)

func StoreAccountDB(account *dto.Account) error {

	db, err := StartDb()
	if err = CheckErr(err); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO account(user, password, cpf, date_joined, balance) values(?,?,?,?,?)")
	if err = CheckErr(err); err != nil {
		return err
	}

	_, err = stmt.Exec(account.Name, account.Secret, account.Cpf, account.Created_at, account.Balance)
	if err = CheckErr(err); err != nil {
		return err
	}

	defer db.Close()
	defer stmt.Close()

	return err
}
