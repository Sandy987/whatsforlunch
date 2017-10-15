package main

import (
	"fmt"
	"net/http"
)

func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to whatsforlunch!")
}

// TODO: Do this better.
var locationHandlers = NewLocationHandlers()

// GetAllRoutes contains all routes
func GetAllRoutes() []*Route {
	return []*Route{
		&Route{
			Name:        "Alive",
			Method:      "GET",
			Pattern:     "alive",
			HandlerFunc: alive,
		},
		&Route{
			Name:        "LocationsList",
			Method:      "GET",
			Pattern:     "location",
			HandlerFunc: locationHandlers.list,
		},
		&Route{
			Name:        "LocationsShow",
			Method:      "GET",
			Pattern:     "location/{locationId}",
			HandlerFunc: locationHandlers.show,
		},
		&Route{
			Name:        "LocationsCreate",
			Method:      "POST",
			Pattern:     "location",
			HandlerFunc: locationHandlers.create,
		},
		&Route{
			Name:        "LocationsPut",
			Method:      "PUT",
			Pattern:     "location",
			HandlerFunc: locationHandlers.update,
		},
	}
}
