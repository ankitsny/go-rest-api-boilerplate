package main

import (
	"fmt"
	"goapi/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// INFO: load app configs here
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// TODO: Connect to database, app level

	// TODO: Connect redis store, app level

	// TODO: Create logger instance? if required

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	log.Printf("server %v is started at %v\n", app.Version, address)
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
