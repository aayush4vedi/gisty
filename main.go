package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// to restrict this root, i.e. not to make it catch-all
	if r.URL.Path != "/" { 
		http.NotFound(w, r) 
		return 
	} 
	w.Write([]byte("Hello from Gisty"))
}

func showGist(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific gist..."))
}

func createGist(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new gist..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)

	log.Println("I love you 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
