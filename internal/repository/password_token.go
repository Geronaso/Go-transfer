package repository

import "fmt"

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

func StoreTokenDB(token string, cpf string) error {
	db, err := StartDb()
	CheckErr(err)

	stmt, err := db.Prepare("INSERT INTO token(cpf, token) values(?,?)")
	CheckErr(err)

	fmt.Println(token)

	_, err = stmt.Exec(cpf, token)
	CheckErr(err)

	defer db.Close()
	defer stmt.Close()

	return err

}
