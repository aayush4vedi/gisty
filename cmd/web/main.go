package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions" // New import

	"github.com/gisty/pkg/models/mysql"
)

type App struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	gists         *mysql.GistModel
	templateCache map[string]*template.Template
	session       *sessions.Session
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "root:password@/gisty?parseTime=true", "DSN for db")

	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "32 bit long secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true // Set the Secure flag on our session cookies

	app := &App{
		errorLog:      errorLog,
		infoLog:       infoLog,
		gists:         &mysql.GistModel{DB: db},
		templateCache: templateCache,
		session:       session,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
		//Add timeouts
		IdleTimeout:  time.Minute, 
        ReadTimeout:  5 * time.Second, 
        WriteTimeout: 10 * time.Second, 
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
