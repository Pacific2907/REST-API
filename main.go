//Basic of Rest API

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Eid      string `json:"empid,omitempty"`
	Name     string `json:"emname,omitempty"`
	Email    string `json:"ememail,omitempty"`
	PContact int    `json:"contact,omitempty"`
}
type Employees []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employees := Employees{
		Employee{Eid: "101", Name: "Prasant chetri", Email: "123@gmail.com"},
	}
	fmt.Println("Your are Successfull")
	json.NewEncoder(w).Encode(employees)
}
func addEmployees(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "POST Request")
}
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT request")
}
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE request")
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees", addEmployees).Methods("POST")
	r.HandleFunc("/employee/{Eid}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{Eid}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}
func main() {
	handleRequests()
}
