package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBCOnnection struct {
	DB *sql.DB
}

var connected = DBCOnnection{}

func Connection() error {
	db, err := sql.Open("postgres", "user=gotest dbname=gotest password=pass#123 sslmode=disable")
	if err != nil {
		return err
	}

	connected = DBCOnnection{
		DB: db,
	}

	return nil
}
