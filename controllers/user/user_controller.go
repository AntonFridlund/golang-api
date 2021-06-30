package user

import (
	"fmt"
	"net/http"
)

func User(res http.ResponseWriter, req *http.Request) {
	fmt.Println("user page")
	fmt.Fprintf(res, "Route: %v\n", req.URL.Path)
}

func NewUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("new user")
	fmt.Fprintf(res, "Route: %v\n", req.URL.Path)
}
