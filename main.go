package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Nigger")
	sql, _ := sql.Open("postgres", "user=postgres dbname=artem sslmode=disable")
	sql.Exec("INSERT INTO users name $1", "artem")
}
