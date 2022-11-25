package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oluwajueda/bookstore-go/pkg/routes"
)

func main (){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:9010", r))
}

