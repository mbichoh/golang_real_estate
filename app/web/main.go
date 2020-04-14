package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//aplication struct to hold the application wide dependencies for the web app
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//add command-line flag to network address
	addr := flag.String("addr", ":4047", "Http Network Address")

	flag.Parse()

	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  //infomation log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) //errors log


	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/estate", app.showEstate)
	mux.HandleFunc("/estate/create", app.createEstate)

	//adding static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//initialize a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Go development server started: http://localhost%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
