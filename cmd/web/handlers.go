package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// to restrict this root, i.e. not to make it catch-all
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. Note that the 
	// home.page.tmpl file must be the *first* file in the slice. 
	files := []string{ 
		"./ui/html/home.page.tmpl", 
		"./ui/html/base.layout.tmpl", 
		"./ui/html/footer.partial.tmpl",
	} 
	// Use the template.ParseFiles() function to read the files and store the 
	// templates in a template set
	ts, err := template.ParseFiles(files...) 
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	//write the template content as the response body.`nil` is in place for dynamic data part
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific gist with ID %d..", id)
}

func createGist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new gist..."))
}
