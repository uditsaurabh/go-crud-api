package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&emp)
	fmt.Println(emp)
	DB.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetALLEmployee(w http.ResponseWriter, r *http.Request) {
	// Get all records
	w.Header().Set("Content-Type", "application/json")
	var emp []Employee
	DB.Find(&emp)
	json.NewEncoder(w).Encode(emp)
	// SELECT * FROM users;
	//json.NewEncoder(w).Encode(result)

}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.First(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(emp)
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.First(&emp, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Save(emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("Employee is deleted")
}
