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
	//restrict access
	mux.Get("/gist/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createGistForm))
	mux.Post("/gist/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createGist))
	
	mux.Get("/gist/:id", dynamicMiddleware.ThenFunc(app.showGist))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm)) 
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser)) 
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm)) 
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser)) 
	//restrict access
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser)) 

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	
	return standardMiddleware.Then(mux)
}
