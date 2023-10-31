package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("mensaje")
}

func RegisterRoutes(app *APP, r *mux.Router) {
	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)
	r.Use(checkMiddlewate)
	r.HandleFunc("/cat", app.GetCatFact)
	r.HandleFunc("/random", app.GetRandomUser)
	r.HandleFunc("/weather", app.Weather)
}
