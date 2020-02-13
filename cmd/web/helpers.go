package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *App) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (app *App) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.AuthenticatedUser = app.authenticatedUser(r) 
	td.CurrentYear = time.Now().Year() 
	td.Flash = app.session.PopString(r, "flash") 
	return td 
}


// The authenticatedUser method returns the ID of the current user from the 
// session, or zero if the request is from an unauthenticated user. 
func (app *App) authenticatedUser(r *http.Request) int { 
	return app.session.GetInt(r, "userID") 
}