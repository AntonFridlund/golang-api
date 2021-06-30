package home

import (
	"fmt"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("home page")
	fmt.Fprintf(res, "Route: %v\n", req.URL.Path)
}
