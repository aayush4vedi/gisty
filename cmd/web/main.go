package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gisty/pkg/models/mysql" // New import
)

// Add a gists field to the application struct. This will allow us to
// make the GistModel object available to our handlers.
type App struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	gists    *mysql.GistModel
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "root:password@/gisty?parseTime=true", "DSN for db")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// pass DSN to openDB
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	//Dont forget to close it
	defer db.Close()

	app := &App{
		errorLog: errorLog,
		infoLog:  infoLog,
		gists: &mysql.GistModel{DB: db}, 
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
