package main

import (
	"log"
	"net/http"
)

//1. Handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {

	//2. Servemux: initialtion
	mux := http.NewServeMux()
	//2. Servemux: Route registration
	mux.HandleFunc("/", home)

	//3. Webserver: start a new web server..
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
