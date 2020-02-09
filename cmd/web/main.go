package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gist/create", createGist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{ 
		Addr:     *addr, 
		ErrorLog: errorLog, 
		Handler:  mux, 
	} 

	infoLog.Printf("I love you %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err) 

}
