package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mbichoh/real_estate/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) //using the notFound helper
		return
	}

	e, err := app.estates.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Estates: e,
	})

}

func (app *application) showEstate(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	e, err := app.estates.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Estate: e,
	})
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
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/estate?id=%d", id), http.StatusSeeOther)

}
