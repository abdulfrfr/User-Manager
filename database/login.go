package database

import (
	"log"
	"net/http"
)

func Login(res http.ResponseWriter, username string) (interface{}, error) {
	var user Users
	rows, err := connected.DB.Query("SELECT * FROM your_table WHERE name = $1", username)

	if err != nil {
		return nil, err
	}

	err = rows.Scan(&user.ID, &user.Name)

	if err != nil {
		log.Fatal(err)
	}

	return user.Name, nil

}
