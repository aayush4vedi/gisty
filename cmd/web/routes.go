package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)

func (app *App) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/gist/create", dynamicMiddleware.ThenFunc(app.createGistForm))
	mux.Post("/gist/create", dynamicMiddleware.ThenFunc(app.createGist))
	mux.Get("/gist/:id", dynamicMiddleware.ThenFunc(app.showGist))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	
	return standardMiddleware.Then(mux)
}
