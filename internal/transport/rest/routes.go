package rest

import "github.com/gorilla/mux"

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	chartas := router.PathPrefix("/chartas").Subrouter()
	withId := chartas.PathPrefix("/{id}").Subrouter()

	chartas.
		Queries(
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(h.CreateNewCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:[\\w\\-]+}",
			"y", "{y:[\\w+\\-]+}",
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(h.SaveRestoredFragmentOfCharta).
		Methods("POST")

	withId.
		Queries(
			"x", "{x:[\\w+\\-]+}",
			"y", "{y:[\\w+\\-]+}",
			"width", "{width:\\w+}",
			"height", "{height:\\w+}",
		).
		HandlerFunc(h.GetPartOfCharta).
		Methods("GET")

	withId.
		HandleFunc("/", h.DeleteCharta).
		Methods("DELETE")

	return router
}
