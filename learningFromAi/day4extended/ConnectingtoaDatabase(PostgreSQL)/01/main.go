package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // <-- Make sure this is included
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "puspo"
	password = "yourpassword"
	dbname   = "testdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	fmt.Println("Successfully connected to the database")
}
