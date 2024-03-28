package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {

	CarregarDados()

	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	http.Handle("/", router)

	log.Println("Server is running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func CarregarDados() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para HealthCheck")
	json.NewEncoder(w).Encode("API está no ar!")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para GetPeople")
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para GetPerson")
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para CreatePerson")
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para UpdatePerson")

	params := mux.Vars(r)
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)

	for index, item := range people {
		if item.ID == params["id"] {
			people[index].Firstname = person.Firstname
			people[index].Lastname = person.Lastname
			people[index].Address = person.Address
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição para DeletePerson")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
