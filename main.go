package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	Email      string `json:"email"`
	Last_name  string `json:"last_name"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Gender     string `json:"gender"`
	Birth_date string `json:"birth_date"`
}

var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if item.Last_name == params["last_name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Last_name = strconv.Itoa(rand.Intn(1000000))
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.Last_name == params["last_name"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.Last_name = params["last_name"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.Last_name == params["last_name"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func main() {
	r := mux.NewRouter()
	persons = append(persons, Person{Last_name: "Devonport", Email: "Logan_Devonport3313@tonsy.org", Country: "Oman", City: "Madrid", Gender: "Male", Birth_date: "Friday, April 4, 8527 8:45 AM"})
	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons/{last_name}", getPerson).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/{last_name}", updatePerson).Methods("PUT")
	r.HandleFunc("/persons/{last_name}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8006", r))
}
