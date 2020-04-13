package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/estate", showEstate)
	mux.HandleFunc("/estate/create", createEstate)

	log.Println("Server starting at port: 4047")
	err := http.ListenAndServe(":4047", mux)
	log.Fatal(err)

}
