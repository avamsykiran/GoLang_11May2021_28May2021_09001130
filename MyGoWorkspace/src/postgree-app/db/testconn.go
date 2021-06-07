package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func VerifyConnection() {

	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to postgre successfully")
}
