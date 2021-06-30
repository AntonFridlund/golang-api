package routes

import (
	controller "main/controllers/home"

	"github.com/gorilla/mux"
)

func Main(router *mux.Router) {
	router.HandleFunc("/", controller.Home).Methods("GET")
	User(router.PathPrefix("/user").Subrouter())
}
