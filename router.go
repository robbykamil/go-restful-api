package main

//imports and package
import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// ensures that routes are matched regardless of trailing slashes.
	// Example: /routes and /routes/ will be treated as the same route
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			  Method(route.Method).
			  Path(route.Pattern).
			  Name(route.Name).
			  Handler(handler)
	}

	return router
}

