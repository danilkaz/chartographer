package rest

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	chartas := router.PathPrefix("/chartas").Subrouter()
	withId := chartas.PathPrefix("/{id}").Subrouter()

	chartas.
		Queries(
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(CreateNewCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:\\w+}",
			"y", "{y:\\w+}",
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(SaveRestoredFragmentOfCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:\\w+}",
			"y", "{y:\\w+}",
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(GetPartOfCharta).
		Methods("GET")

	withId.
		HandleFunc("/", DeleteCharta).
		Methods("DELETE")

	return router
}
