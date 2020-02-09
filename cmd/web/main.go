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

	fileServer := http.FileServer(http.Dir("./ui/static/")) 
	//For matching paths, we strip the"/static" prefix before the request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) 

	log.Println("I love you 3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
