package rest

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	chartas := router.PathPrefix("/chartas").Subrouter()
	withId := chartas.PathPrefix("/{id}").Subrouter()

	chartas.
		Queries(
			"width", "{width:[0-9]+}",
			"height", "{height:[0-9]+}",
		).
		HandlerFunc(CreateNewCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:[0-9]+}",
			"y", "{y:[0-9]+}",
			"width", "{width:[0-9]+}",
			"height", "{height:[0-9]+}",
		).
		HandlerFunc(SaveRestoredFragmentOfCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:[0-9]+}",
			"y", "{y:[0-9]+}",
			"width", "{width:[0-9]+}",
			"height", "{height:[0-9]+}",
		).
		HandlerFunc(GetPartOfCharta).
		Methods("GET")

	withId.
		HandleFunc("/", DeleteCharta).
		Methods("DELETE")

	return router
}
