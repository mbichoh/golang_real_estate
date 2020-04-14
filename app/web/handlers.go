package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) //using the notFound helper
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/_footer.partial.tmpl",
		"./ui/html/_header.partial.tmpl",
		"./ui/html/_javascript.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) //using serverError helper
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) //using serverError helper
	}
}

func (app *application) showEstate(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "The Estate in page ID %d", id)
}

func (app *application) createEstate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // Using the clientError() helper.
		return
	}
	w.Write([]byte("Create new estate"))
}
