package database

import (
	"encoding/json"
	"log"
	"net/http"
)

// responsible for handling response for all requests

func Res[U []Users | Users | map[string]string](res http.ResponseWriter, result U) {

	res.Header().Set("Content-type", "application/json")

	var final []byte
	var err error
	final, err = json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	res.WriteHeader(http.StatusOK)
	res.Write(final)

}

// error manager for errors from the queries
func errorRes(res http.ResponseWriter, err error) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(err.Error()))

}
