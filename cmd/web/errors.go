package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//sends a generic 500 Internal Server Error response to the user.
func (app *App) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//sends a specific status code and corresponding description
func (app *App) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//sends a 404 Not Found
func (app *App) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
