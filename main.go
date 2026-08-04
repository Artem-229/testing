package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	sql, err := sql.Open("postgres", "user=postgres dbname=artem sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec("INSERT INTO users name $1", "artem")

	run()
}

func run() error {
	return fmt.Errorf("test error")
}
