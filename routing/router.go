package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route describes the components of an HTTP route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains a slice of Route
type Routes []Route

// NewRouter returns a new, logged HTTP router that acts over the given routes
func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(fmt.Sprintf("/api/%s", route.Pattern)).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
