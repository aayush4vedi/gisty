package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)

func (app *App) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/gist/create", http.HandlerFunc(app.createGistForm))
	mux.Post("/gist/create", http.HandlerFunc(app.createGist))
	mux.Get("/gist/:id", http.HandlerFunc(app.showGist))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	
	return standardMiddleware.Then(mux)
}
