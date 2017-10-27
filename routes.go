package main

import (
	"fmt"
	"net/http"
)

func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to whatsforlunch!")
}

// TODO: Do this better?
var locationHandlers = NewLocationHandlers()
var dishHandlers = NewDishHandlers()

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
		&Route{
			Name:        "DishesList",
			Method:      "GET",
			Pattern:     "dish",
			HandlerFunc: dishHandlers.list,
		},
		&Route{
			Name:        "DishesShow",
			Method:      "GET",
			Pattern:     "dish/{dishId}",
			HandlerFunc: dishHandlers.show,
		},
		&Route{
			Name:        "DishesCreate",
			Method:      "POST",
			Pattern:     "dish",
			HandlerFunc: dishHandlers.create,
		},
		&Route{
			Name:        "DishesPut",
			Method:      "PUT",
			Pattern:     "dish",
			HandlerFunc: dishHandlers.update,
		},
	}
}
