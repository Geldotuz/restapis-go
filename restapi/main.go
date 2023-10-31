package main

import (
	"log"
	"net/http"
	"routes-api/rest"

	"github.com/gorilla/mux"
)

func main() {
	app := &rest.APP{}
	r := mux.NewRouter()
	rest.RegisterRoutes(app, r)

	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	log.Println("listening")
	srv.ListenAndServe()
}
