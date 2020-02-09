package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)

	log.Println("I love you 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
