package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gisty/pkg/forms"
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

	// data this will return the empty string.
	flash := app.session.PopString(r, "flash")

	app.render(w, r, "show.page.tmpl", &templateData{
		Flash: flash,
		Gist:  s,
	})
}

func (app *App) createGist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := app.gists.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Gist successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/gist/%d", id), http.StatusSeeOther)
}

func (app *App) createGistForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *App) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *App) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 4)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail { 
		form.Errors.Add("email", "Address is already in use") 
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form}) 
		return 
	} else if err != nil { 
		app.serverError(w, err) 
		return 
	} 
	app.session.Put(r, "flash", "Your signup was successful. Please log in.") 
	http.Redirect(w, r, "/user/login", http.StatusSeeOther) 
}

func (app *App) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{ 
		Form: forms.New(nil), 
	}) 
}
func (app *App) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() 
	if err != nil { 
		app.clientError(w, http.StatusBadRequest) 
		return 
	} 
	form := forms.New(r.PostForm) 
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password")) 
	if err == models.ErrInvalidCredentials { 
		form.Errors.Add("generic", "Email or Password is incorrect") 
		app.render(w, r, "login.page.tmpl", &templateData{Form: form}) 
		return 
	} else if err != nil { 
		app.serverError(w, err) 
		return 
	} 
	app.session.Put(r, "userID", id) 
	http.Redirect(w, r, "/gist/create", http.StatusSeeOther) 
}

func (app *App) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "userID") 
	app.session.Put(r, "flash", "You've been logged out successfully!") 
	http.Redirect(w, r, "/", 303) 
}

// Testing function
func ping(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte("OK")) 
} 