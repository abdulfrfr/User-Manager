package database

import (
	"encoding/json"
	"log"
	"net/http"
)

func Res(res http.ResponseWriter, result []Users) {

	res.Header().Set("Content-type", "application/json")

	res.WriteHeader(http.StatusOK)

	var final []byte
	var err error
	final, err = json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	res.Write(final)

}
