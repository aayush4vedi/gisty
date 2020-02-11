package main

import "net/http"

// Update the signature for the routes() method so that it returns a http.Handler
func (app *App) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return secureHeaders(mux)
}
