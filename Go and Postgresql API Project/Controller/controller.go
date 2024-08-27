package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/igauravrana/Models"
	connection "github.com/igauravrana/PostgresConnection"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//use connection.DB to setup the connection first it will set connection with models folder func SeeAllData
	employees, err := models.SeeAllData(connection.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//grab the request from the response body
	id := mux.Vars(r)["id"]

	// convert the response into int value
	conv, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	data := models.ReadEmployees(connection.DB, conv)
	json.NewEncoder(w).Encode(data)
}

func CreateOneEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var data models.PostgresStruct
	json.NewDecoder(r.Body).Decode(&data)
	models.CreateEmployee(connection.DB, data)
	json.NewEncoder(w).Encode(data)
}

func UpdateOneEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	id := mux.Vars(r)["id"]

	conv, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	var data models.PostgresStruct

	json.NewDecoder(r.Body).Decode(&data)
	data.Employee_id = conv

	models.UpdateEmployee(connection.DB, data)
	json.NewEncoder(w).Encode(data)
}

func DeleteOneEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	id := mux.Vars(r)["id"]

	conv, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	var data models.PostgresStruct

	models.DeleteEmployee(connection.DB, conv)
	json.NewEncoder(w).Encode(data)
}
