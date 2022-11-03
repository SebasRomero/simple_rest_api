package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet { //This shows when the method used was different to GET
		w.WriteHeader(http.StatusMethodNotAllowed) //This show us the corresponding method header
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	fmt.Fprintf(w, "Hello there %s", "visitor")
}

//Get a country from the array in json format
func getCountry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

//Add a country in the array
func addCountry(w http.ResponseWriter, r *http.Request) {

	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, country)
	fmt.Fprintf(w, "Country was added")
}
