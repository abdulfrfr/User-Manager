package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abdulfrfr/user-manager/database"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(res http.ResponseWriter, req *http.Request) {

	var credential Credentials

	err := json.NewDecoder(req.Body).Decode(&credential)

	if err != nil {
		log.Fatal(err)
	}

	database.Login(res)

}
