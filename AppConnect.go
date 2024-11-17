package main

import (
	"log"
	"net/http"

	"github.com/abdulfrfr/user-manager/database"
	"github.com/abdulfrfr/user-manager/users"
)

// function for database and http connections, semi-entry point for our application.

func connect(address string) {

	http.HandleFunc("GET /", users.GetUsers)

	err := database.Connection()

	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}

}
