package users

import (
	"log"
	"net/http"

	"github.com/user-manager/database"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")

	err := database.GetUsers(res)

	if err != nil {
		log.Fatal(err)
	}

}
