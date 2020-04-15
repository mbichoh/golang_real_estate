package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

//500 internal server error
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//400 bad request
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//404 page not found
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

//render template from cache to avoid duplicate code
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData){
	ts, ok := app.templateCache[name]
	if !ok{
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil{
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}