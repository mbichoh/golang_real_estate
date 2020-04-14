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

	agentID := 1
	name := "Nyumba Itu"
	address := "Baricho"
	county := "Kirinyaga"
	shortDesc := "Pretty nice home"
	longDesc := "As the description, You will never get such an awesome, amaizing place. Its dope and fashioned just as you like"
	bedroom := 2
	washroom := 2
	spaceArea := 1290
	packing := 1
	price := 120000.00

	id, err := app.estates.Insert(agentID, name, address, county, shortDesc, longDesc, bedroom, washroom, spaceArea, packing, price)
	if err != nil{
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/estate?id=%d", id), http.StatusSeeOther)

}
