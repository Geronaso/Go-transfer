package repository

import (
	"go-transfer/internal/datastruct"
)

func RetrieveAccount(cpf string) (*datastruct.Account, error) {
	db, err := StartDb()
	if err = CheckErr(err); err != nil {
		return nil, err
	}

	query, err := db.Query("SELECT * FROM account WHERE cpf=?", cpf)
	if err = CheckErr(err); err != nil {
		return nil, err
	}

	var place_holder string
	account := new(datastruct.Account)
	query.Next()
	err = query.Scan(&account.Id, &account.Name, &account.Date_Joined, &account.Cpf, &account.Balance, &place_holder)
	if err = CheckErr(err); err != nil {
		return nil, err
	}

	defer query.Close()
	defer db.Close()

	return account, err

}

func TransferDB(accounts *datastruct.TransferValues, transfer *datastruct.Transfer) error {

	db, err := StartDb()
	if err = CheckErr(err); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO transfers(account_origin_id, account_destination_id, amount, created_at) values(?,?,?,?)")
	if err = CheckErr(err); err != nil {
		return err
	}

	_, err = stmt.Exec(transfer.Account_origin_id, transfer.Account_destination_id, transfer.Amount, transfer.Created_at)
	if err = CheckErr(err); err != nil {
		return err
	}

	stmt, err = db.Prepare("UPDATE account SET balance = ? WHERE id = ?")
	if err = CheckErr(err); err != nil {
		return err
	}

	_, err = stmt.Exec(accounts.Origin.Balance, accounts.Origin.Id)
	if err = CheckErr(err); err != nil {
		return err
	}

	stmt, err = db.Prepare("UPDATE account SET balance = ? WHERE id = ?")
	if err = CheckErr(err); err != nil {
		return err
	}

	_, err = stmt.Exec(accounts.Destination.Balance, accounts.Destination.Id)
	if err = CheckErr(err); err != nil {
		return err
	}

	defer stmt.Close()
	defer db.Close()

	return err

}

func RetrieveTransferDB(account *datastruct.Account) ([]datastruct.Transfer, error) {
	db, err := StartDb()
	if err = CheckErr(err); err != nil {
		return nil, err
	}

	query, err := db.Query("SELECT * FROM transfers WHERE account_origin_id = ?", account.Id)
	if err = CheckErr(err); err != nil {
		return nil, err
	}

	transfers_info := make([]datastruct.Transfer, 0)

	for query.Next() {
		current_transfer := datastruct.Transfer{}
		err = query.Scan(&current_transfer.Id, &current_transfer.Account_origin_id, &current_transfer.Account_destination_id, &current_transfer.Amount, &current_transfer.Created_at)
		if err = CheckErr(err); err != nil {
			return nil, err
		}

		transfers_info = append(transfers_info, current_transfer)
	}

	defer db.Close()
	defer query.Close()

	return transfers_info, err
}
