package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/gisty/pkg/forms"
	"github.com/gisty/pkg/models"
)

type templateData struct {
	Gist        *models.Gist
	Gists       []*models.Gist
	CurrentYear int
	Form        *forms.Form
	Flash       string 
	AuthenticatedUser int  // new field
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var datefunc = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(datefunc).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
