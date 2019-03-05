package main

import "github.com/gorilla/mux"

// App ...
type App struct {
	Router *mux.Router
}

// Initialize ...
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}

// Run ...
func (a *App) Run(addr string) {

}
