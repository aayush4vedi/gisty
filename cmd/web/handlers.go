package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gisty/pkg/models"
)

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Gists: s,
	})
}

func (app *App) showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.gists.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "show.page.tmpl", &templateData{
		Gist: s,
	})
}

func (app *App) createGist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() 
	if err != nil { 
		app.clientError(w, http.StatusBadRequest) 
		return 
	} 
	title := r.PostForm.Get("title") 
	content := r.PostForm.Get("content") 
	expires := r.PostForm.Get("expires") 
	id, err := app.gists.Insert(title, content, expires) 
	if err != nil { 
		app.serverError(w, err) 
		return 
	} 
	http.Redirect(w, r, fmt.Sprintf("/gist/%d", id), http.StatusSeeOther)
}

func (app *App) createGistForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)  
}