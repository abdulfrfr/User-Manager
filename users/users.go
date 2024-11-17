package users

import (
	"log"
	"net/http"

	"github.com/abdulfrfr/user-manager/database"
)

// uses the GetUsers function from the database package.

func GetUsers(res http.ResponseWriter, req *http.Request) {

	err := database.GetUsers(res)

	if err != nil {
		log.Fatal(err)
	}

}
