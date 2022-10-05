package repository

import (
	"go-transfer/internal/dto"
)

func StoreAccountDB(account *dto.Account) error {

	db, err := StartDb()
	CheckErr(err)

	stmt, err := db.Prepare("INSERT INTO account(user, password, cpf, date_joined, balance) values(?,?,?,?,?)")
	CheckErr(err)

	_, err = stmt.Exec(account.Name, account.Secret, account.Cpf, account.Created_at, account.Balance)
	CheckErr(err)

	defer db.Close()
	defer stmt.Close()

	return err
}
