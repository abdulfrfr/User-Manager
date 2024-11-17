package database

import (
	"net/http"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// function to query database to get all rows in the list table serves as a semi handler for our GetUser handle function, makes use of the database connection object from the DatabaseConnection.go file
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
			errorRes(res, err)
		}

		allUsers = append(allUsers, singleUser)

	}

	Res(res, allUsers)
	return nil
}
