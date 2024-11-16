package users

import (
	"log"
	"net/http"

	"github.com/abdulfrfr/user-manager/database"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	err := database.GetUsers(res)

	if err != nil {
		log.Fatal(err)
	}

}
