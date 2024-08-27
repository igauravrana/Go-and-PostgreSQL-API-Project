package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var DB *sql.DB

func Connect() {

	//Paste your PostgreSql database address here
	conStr := "postgres://user:password@localhost:5432/Project%201?sslmode=disable"

	DB, _ = sql.Open("postgres", conStr)

	//check ping between postgres and go
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection is established")
}
