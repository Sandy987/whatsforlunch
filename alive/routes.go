package alive

import (
	"fmt"
	"net/http"

	"github.com/Sandy987/whatsforlunch/routing"
)

func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to whatsforlunch!")
}

// Routes contains all routes for the alive status
var Routes = routing.Routes{
	routing.Route{
		Name:        "Alive",
		Method:      "GET",
		Pattern:     "alive",
		HandlerFunc: alive,
	},
}
