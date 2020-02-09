package main

import (
	"fmt"
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
	w.Write([]byte("Hello from Gisty"))
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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)

	log.Println("I love you 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
