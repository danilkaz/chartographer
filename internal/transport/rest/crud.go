package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) CreateNewCharta(w http.ResponseWriter, r *http.Request) {
	var width, height int
	vars := mux.Vars(r)
	values := []string{vars["width"], vars["height"]}
	if err := validateValuesAndUnpack(values, &width, &height); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.services.Create(width, height)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, id)
}

func (h *Handler) SaveRestoredFragmentOfCharta(w http.ResponseWriter, r *http.Request) {
	var x, y, width, height int
	vars := mux.Vars(r)
	values := []string{vars["width"], vars["height"], vars["x"], vars["y"]}
	if err := validateValuesAndUnpack(values, &width, &height, &x, &y); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetPartOfCharta(w http.ResponseWriter, r *http.Request) {
	var x, y, width, height int
	vars := mux.Vars(r)
	values := []string{vars["width"], vars["height"], vars["x"], vars["y"]}
	if err := validateValuesAndUnpack(values, &width, &height, &x, &y); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handler) DeleteCharta(w http.ResponseWriter, r *http.Request) {
}
