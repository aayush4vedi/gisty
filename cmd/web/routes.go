package main

import (
	"net/http"

	"github.com/justinas/alice" // New import
)

func (app *App) routes() http.Handler {

	// Create a middleware chain containing our 'standard' middleware 
	// which will be used for every request our application receives. 
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
		
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Return the 'standard' middleware chain followed by the servemux.
	return standardMiddleware.Then(mux)
}
