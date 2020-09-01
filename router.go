package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, article := range articles {
		var handler http.Handler

		handler = article.HandlerFunc
		handler = Logger(handler, article.Name)

		router.
			Methods(article.Method).
			Path(article.Pattern).
			Name(article.Name).
			Handler(handler)

	}

	return router
}
