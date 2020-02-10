package main

import "net/http"

func (app *App) routes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}