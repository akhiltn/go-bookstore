package main

import (
	"log"
	"net/http"

	"github.com/akhiltn/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Starting server on port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
