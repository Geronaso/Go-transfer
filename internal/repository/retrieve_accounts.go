package repository

import "go-transfer/internal/datastruct"

func GetAccountsDB() ([]datastruct.Account, error) {

	db, err := StartDb()
	CheckErr(err)

	query, err := db.Query("SELECT * FROM account")
	CheckErr(err)

	accounts_info := make([]datastruct.Account, 0)

	for query.Next() {
		current_account := datastruct.Account{}
		var place_holder string
		err = query.Scan(&current_account.Id, &current_account.Name, &current_account.Date_Joined, &current_account.Cpf, &current_account.Balance, &place_holder)
		CheckErr(err)

		accounts_info = append(accounts_info, current_account)
	}

	defer db.Close()

	return accounts_info, err
}

func GetBalanceDB(user string) (string, error) {

	db, err := StartDb()
	CheckErr(err)

	query, err := db.Query("SELECT balance FROM account WHERE user=?", user)
	CheckErr(err)

	var balance string
	query.Next()
	err = query.Scan(&balance)

	defer db.Close()

	return balance, err
}
