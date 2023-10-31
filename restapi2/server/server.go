package server

import "net/http"

type Country struct {
	Name     string
	Language string
}

var countries = []Country{}

func New(addr string) *http.Server { // recibir diferentes request addr = ":8080"
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
