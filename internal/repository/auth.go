package repository

func RetrieveHashPass(user string) (string, error) {

	db, err := StartDb()
	CheckErr(err)

	query, err := db.Query("SELECT password FROM account WHERE cpf=?", user)
	CheckErr(err)

	var hashed_pass string
	query.Next()
	err = query.Scan(&hashed_pass)

	defer query.Close()
	defer db.Close()

	return hashed_pass, err

}
