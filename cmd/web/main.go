package main

import (
	"database/sql"
	"flag"
	"html/template"  // New import 
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gisty/pkg/models/mysql"
)

type App struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	gists         *mysql.GistModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "root:password@/gisty?parseTime=true", "DSN for db")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new template cache... 
	templateCache, err := newTemplateCache("./ui/html/") 
	if err != nil { 
		errorLog.Fatal(err) 
	} 

	// And add it to the application dependencies. 
	app := &App{
		errorLog: errorLog,
		infoLog:  infoLog,
		gists:    &mysql.GistModel{DB: db},
		templateCache: templateCache, 
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server on: %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
