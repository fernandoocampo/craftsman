package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartWebServer starts a new http server.
func StartWebServer(port string) {
	http.Handle("/", NewRouter())

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Panicf("An error occured starting HTTP listener at port %s, error %s", port, err)
	}
}

// NewRouter returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {

	// Create an instance of the Gorilla router
	// Gorilla router matches incoming requests against a list of
	// registered routes and calls a handler for the route that matches
	// the URL or other conditions
	router := mux.NewRouter().StrictSlash(true)

	// Get a craftsman by its id
	router.Methods("GET").
		Path("/{id}").
		Name("findByID").
		HandlerFunc(findByID)

	// create a craftsman
	router.Methods("POST").
		Path("/").
		Name("create").
		HandlerFunc(create)

	// update a craftsman
	router.Methods("PUT").
		Path("/").
		Name("update").
		HandlerFunc(update)

	// update a craftsman's state
	router.Methods("PUT").
		Path("/{id}/state").
		Name("updateState").
		HandlerFunc(updateState)

	// get health status of this service.
	router.Methods("GET").
		Path("/health").
		Name("health").
		HandlerFunc(health) // what's the health

	return router

}
