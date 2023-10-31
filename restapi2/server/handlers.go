package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// diferentes funciones en cada ruta

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	if r.Method != http.MethodGet { // m1
		w.WriteHeader(http.StatusMethodNotAllowed) // StatusCode
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Welcome %s", "Visitor")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func addCountry(w http.ResponseWriter, r *http.Request) { //endpoin

	country := Country{}

	fmt.Print(country)

	err := json.NewDecoder(r.Body).Decode(&country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Print(country)

	countries = append(countries, country)
	fmt.Fprintf(w, "country was added")
}
