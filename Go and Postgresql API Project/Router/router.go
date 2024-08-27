package router

import (
	"fmt"

	"github.com/gorilla/mux"
	controller "github.com/igauravrana/Controller"
)

func Router() *mux.Router {
	fmt.Println("Welcome to PostgreSql api")

	//make a new router
	r := mux.NewRouter()

	//handles the function of r router
	r.HandleFunc("/getall", controller.GetAllEmployees).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", controller.GetEmployees).Methods("GET")
	r.HandleFunc("/create", controller.CreateOneEmployee).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", controller.UpdateOneEmployee).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", controller.DeleteOneEmployee).Methods("DELETE")
	return r
}
