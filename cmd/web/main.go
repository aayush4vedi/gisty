package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//Define an application struct to hold the application-wide dependencies for the app & make it available for handler
type application struct { 
	errorLog *log.Logger 
	infoLog  *log.Logger 
} 
	
func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Initialize a new instance of application containing the dependencies. 
	app := &application{ 
		errorLog: errorLog, 
		infoLog:  infoLog, 
	} 
	
	mux := http.NewServeMux()

	// Swap the route declarations to use the application struct's methods as handler functions
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist", app.showGist)
	mux.HandleFunc("/gist/create", app.createGist)

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
