package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
)

//aplication struct to hold the application wide dependencies for the web app
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//add command-line flag to network address
	addr := flag.String("addr", ":4047", "Http Network Address")
	dsn := flag.String("dsn", "root:@/real_estate?parseTime=true", "Mysql data source name")
	flag.Parse()

	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  //infomation log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) //errors log

	//passing openDb to dsn from the command line
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	//close db before main() function exits
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//initialize a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), //calling the app.routes() method
	}

	infoLog.Printf("Go development server started: http://localhost%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
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
