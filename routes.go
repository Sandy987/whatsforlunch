package main

import (
	"fmt"
	"net/http"
)

func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to whatsforlunch!")
}

// Routes contains all routes for the alive status
func GetAllRoutes() []*Route {
	return []*Route{
		&Route{
			Name:        "Alive",
			Method:      "GET",
			Pattern:     "alive",
			HandlerFunc: alive,
		},
	}
}
