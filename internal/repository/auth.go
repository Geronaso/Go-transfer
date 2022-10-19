package repository

func RetrieveHashPassDB(user string) (string, error) {

	query, err := db.Query("SELECT password FROM account WHERE cpf=?", user)
	if err = CheckErr(err); err != nil {
		return "", err
	}
	defer query.Close()

	var hashed_pass string
	query.Next()
	err = query.Scan(&hashed_pass)

	return hashed_pass, err

}
