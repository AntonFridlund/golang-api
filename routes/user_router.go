package routes

import (
	controller "main/controllers/user"

	"github.com/gorilla/mux"
)

func User(router *mux.Router) {
	router.HandleFunc("", controller.User).Methods("GET")
	router.HandleFunc("/new", controller.NewUser).Methods("GET")
}
