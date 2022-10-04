package repository

import (
	"go-transfer/internal/dto"
)

func StoreAccountDB(p *dto.Account) error {

	db, err := StartDb()
	CheckErr(err)

	stmt, err := db.Prepare("INSERT INTO account(user, password, cpf, date_joined, balance) values(?,?,?,?,?)")
	CheckErr(err)

	_, err = stmt.Exec(p.Name, p.Secret, p.Cpf, p.Created_at, p.Balance)
	CheckErr(err)

	defer db.Close()
	defer stmt.Close()

	return err
}
