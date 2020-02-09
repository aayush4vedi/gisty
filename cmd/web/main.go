package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type App struct { 
	errorLog *log.Logger 
	infoLog  *log.Logger 
} 
	
func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &App{ 
		errorLog: errorLog, 
		infoLog:  infoLog, 
	} 
	
	srv := &http.Server{ 
		Addr:     *addr, 
		ErrorLog: errorLog, 
		Handler:  app.routes(), 
	} 

	infoLog.Printf("I love you %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err) 
}
