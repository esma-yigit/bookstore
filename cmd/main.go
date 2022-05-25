package main

import (
	"github.com/esma-yigit/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	err := http.ListenAndServe("/", r)
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
