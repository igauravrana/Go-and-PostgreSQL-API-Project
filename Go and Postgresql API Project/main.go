package main

import (
	"fmt"
	"log"
	"net/http"

	connection "github.com/igauravrana/PostgresConnection"
	router "github.com/igauravrana/Router"
)

func main() {
	fmt.Println("Welcome to the API made by Gaurav")

	//call the connect func from connection folder to setup the connection between go and postgres
	connection.Connect()

	//listen and serve at port 8000
	log.Fatal(http.ListenAndServe(":8000", router.Router()))
}
