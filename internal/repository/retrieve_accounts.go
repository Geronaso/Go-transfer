package repository

import "go-transfer/internal/datastruct"

func GetAccountsDB() ([]datastruct.Account, error) {

	query, err := db.Query("SELECT * FROM account")
	if err = CheckErr(err); err != nil {
		return nil, err
	}
	defer query.Close()

	accounts_info := make([]datastruct.Account, 0)

	for query.Next() {
		current_account := datastruct.Account{}
		var place_holder string
		err = query.Scan(&current_account.Id, &current_account.Name, &current_account.Created_at, &current_account.Cpf, &current_account.Balance, &place_holder)
		if err = CheckErr(err); err != nil {
			return nil, err
		}

		accounts_info = append(accounts_info, current_account)
	}

	return accounts_info, err
}

func GetBalanceDB(user string) (string, error) {

	query, err := db.Query("SELECT balance FROM account WHERE user=?", user)
	if err = CheckErr(err); err != nil {
		return "", err
	}
	defer query.Close()

	var balance string
	query.Next()
	err = query.Scan(&balance)

	return balance, err
}
