package main

import (
	"log"
	"net/http"

	"github.com/cspark0610/go-mysql-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9900", r))
}
