package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// TODO: load app configs here

	// TODO: Connect to database, app level

	// TODO: Connect redis store, app level

	// TODO: Create logger instance? if required

	// start the server
	address := fmt.Sprintf(":%v", 8080)                    // get the port number from the app config
	log.Printf("server %v is started at %v\n", 1, address) // get the app version from the app config
	panic(http.ListenAndServe(address, buildRoutes()))
}

func buildRoutes() *mux.Router {
	// create instance of mux router
	r := mux.NewRouter()

	// Initialize not found handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Rosource not found"))
	})

	// Set path prefix/route group
	r.PathPrefix("v1")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")

	return r
}
