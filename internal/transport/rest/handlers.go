package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateNewCharta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	fmt.Println("create")
}

func SaveRestoredFragmentOfCharta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	fmt.Println("save")
}

func GetPartOfCharta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	fmt.Println("get")
}

func DeleteCharta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	fmt.Println("delete")
}
