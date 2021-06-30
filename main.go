package main

import (
	"log"
	"main/routes"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Let Main router handle all top level routes.
	routes.Main(router.PathPrefix("/").Subrouter())

	server := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	log.Fatal(server.ListenAndServe())
}
