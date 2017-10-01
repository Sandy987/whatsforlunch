package user

import (
	"github.com/Sandy987/whatsforlunch/routing"
)

// Routes contains all routes for the alive status
var Routes = routing.Routes{
	routing.Route{
		Name:        "Show",
		Method:      "GET",
		Pattern:     "user/{userId}",
		HandlerFunc: Show,
	},
	routing.Route{
		Name:        "Signup",
		Method:      "POST",
		Pattern:     "signup",
		HandlerFunc: Signup,
	},
	routing.Route{
		Name:        "Update",
		Method:      "POST",
		Pattern:     "user/{userId}",
		HandlerFunc: Update,
	},
}
