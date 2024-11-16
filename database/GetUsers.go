package database

import (
	"log"
	"net/http"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUsers(res http.ResponseWriter) error {
	row, err := connected.DB.Query("SELECT * FROM list")

	if err != nil {
		return err
	}

	allUsers := []Users{}
	for row.Next() {
		singleUser := Users{}
		err = row.Scan(&singleUser.ID, &singleUser.Name)

		if err != nil {
			log.Fatal(err)
		}

		allUsers = append(allUsers, singleUser)

	}

	Res(res, allUsers)
	return nil
}
