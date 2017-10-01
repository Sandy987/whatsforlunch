package main

import (
	"log"
	"net/http"

	"github.com/Sandy987/whatsforlunch/alive"
	"github.com/Sandy987/whatsforlunch/domain"
	"github.com/Sandy987/whatsforlunch/routing"
	"github.com/Sandy987/whatsforlunch/user"
	"github.com/gorilla/mux"
)

// App is the container for the whatsforlunch API and contains routing and other
// context information
type App struct {
	Router *mux.Router
}

// Initialize sets up all app state and contexts
func (a *App) Initialize() {
	// TODO: Don't append routes like this?
	allRoutes := append(alive.Routes, user.Routes...)

	a.Router = routing.NewRouter(allRoutes)
}

// StartAPI starts the app on a particular address
func (a *App) StartAPI(addr string) {
	db := domain.InitDb()
	err := domain.MigrateToLatest()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
