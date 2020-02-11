package main

import (
	"fmt"
	"html/template"
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
	// Create an instance of a templateData struct holding the slice of 
	// snippets. 
	data := &templateData{Gists: s} 
	
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Pass in the templateData struct when executing the template
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
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

	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{Gist: s}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, data)  		// Pass in the templateData struct when executing the template
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
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

func (app *App) render(w http.ResponseWriter, r *http.Request, name string, td map[string]*template.Template){
	// Retrieve the appropriate template set from the cache based on the page n
	// (like 'home.page.tmpl'). If no entry exists in the cache with the 
	// provided name, call the serverError helper method that we made earlier. 
	ts, ok := app.templateCache[name] 
	if !ok { 
		app.serverError(w, fmt.Errorf("The template %s does not exist", name)) 
		return 
	}
	// Execute the template set, passing in any dynamic data. 
	err := ts.Execute(w, td) 
	if err != nil { 
		app.serverError(w, err) 
	} 
}