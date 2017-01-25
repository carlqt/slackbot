package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route{
// 	"tweet_tcl",
// 	"POST",
// 	"/tweet_tcl",
// 	Tweet,
// }

// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(TokenHandler(handler, "FQoVHazAecqMVK540tqmLOGg"), route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)

	}
	return router
}
