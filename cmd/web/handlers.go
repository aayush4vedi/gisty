package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gisty/pkg/models"
)

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "home.page.tmpl", &templateData{
		Gists: s,
	})
}

func (app *App) showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	id, err := app.gists.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/gist?id=%d", id), http.StatusSeeOther)
}

func (app *App) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter. If there's an error, call our serverError helper and return.
	err := ts.Execute(w, buf)
	if err != nil {
		app.serverError(w, err)
		return 
	}

	// Write the contents of the buffer to the http.ResponseWriter
	buf.WriteTo(w) 
}
