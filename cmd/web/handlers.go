package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)	// Use the notFound() helper 
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w,err)  // Use the serverError() helper. 
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w,err)  // Use the serverError() helper. 
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *App) showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)	// Use the notFound() helper 
		return
	}
	fmt.Fprintf(w, "Display a specific gist with ID %d..", id)
}

//Methodising the function
func (app *App) createGist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)   //use clientError
		return
	}
	w.Write([]byte("Create a new gist..."))
}
