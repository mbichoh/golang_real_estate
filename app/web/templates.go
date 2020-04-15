package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/mbichoh/real_estate/pkg/models"
)

type templateData struct {
	CurrentYear int
	Estate      *models.Estate
	Estates     []*models.Estate
}

//create a readable humanDate function
func humanDate(t time.Time) string{
	return t.Format("02 Jan 2006 at 15:04")
}

var ownFunctions  = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//getting slice of all filepaths with extension '.page.tmpl'
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	//loop through pages 1 by 1
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file path and assign it to the name variable.
		name := filepath.Base(page)

		//Parse the page template file in to a template set.
		ts, err := template.New(name).Funcs(ownFunctions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//using parseGlob, add any 'layout' templates(our base layout)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		//using parseGlob, add all any 'partials' templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add the template set to the cache, using the name of the page as the key
		cache[name] = ts
	}

	//return the map
	return cache, nil
}
