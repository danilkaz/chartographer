package rest

import (
	"fmt"
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/image/bmp"
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
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := fmt.Fprint(w, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SaveRestoredFragmentOfCharta(w http.ResponseWriter, r *http.Request) {
	var x, y, width, height int
	vars := mux.Vars(r)
	values := []string{vars["width"], vars["height"], vars["x"], vars["y"]}
	_, exists := vars["id"]
	if err := validateValuesAndUnpack(values, &width, &height, &x, &y); err != nil || !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	img, err := bmp.Decode(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.services.SaveRestoredFragment(id, x, y, width, height, models.NewCharta(img))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetPartOfCharta(w http.ResponseWriter, r *http.Request) {
	var x, y, width, height int
	vars := mux.Vars(r)
	values := []string{vars["width"], vars["height"], vars["x"], vars["y"]}
	_, exists := vars["id"]
	if err := validateValuesAndUnpack(values, &width, &height, &x, &y); err != nil || !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	subCharta, err := h.services.GetPart(id, x, y, width, height)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = bmp.Encode(w, subCharta.Image); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteCharta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, exists := vars["id"]; !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = h.services.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
