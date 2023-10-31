package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index) // x

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			{
				getCountries(w, r)
			}
		case http.MethodPost:
			{
				addCountry(w, r)
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Mehod not allowed")
			return
		}

	})
}
