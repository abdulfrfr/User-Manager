package main

import (
	"log"
	"net/http"

	"github.com/user-manager/database"
	"github.com/user-manager/users"
)

func connect(address string) {

	http.HandleFunc("GET /", users.CreateUser)

	err := database.Connection()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(address, nil)

}
