package repository

import (
	"go-transfer/internal/datastruct"
)

func StoreAccountDB(account *datastruct.Account) error {

	stmt, err := db.Prepare("INSERT INTO account(user, password, cpf, date_joined, balance) values(?,?,?,?,?)")
	if err = CheckErr(err); err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.Name, account.Secret, account.Cpf, account.Created_at, account.Balance)
	if err = CheckErr(err); err != nil {
		return err
	}

	return err
}
