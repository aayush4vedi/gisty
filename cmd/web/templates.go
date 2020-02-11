package main

import (
	"html/template"
	"path/filepath"

	"github.com/gisty/pkg/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Gist  *models.Gist
	Gists []*models.Gist
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the filepath.Glob function to get a slice of all filepaths with
	// the extension '.page.tmpl'. This essentially gives us a slice of all the
	// 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add any 'layout' templates to the 
		// template set (in our case, it's just the 'base' layout at the 
		// moment). 
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl")) 
		if err != nil { 
			return nil, err 
		} 
		// Use the ParseGlob method to add any 'partial' templates to the 
		// template set (in our case, it's just the 'footer' partial at the 
		// moment). 
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl")) 
		if err != nil { 
			return nil, err 
		}
		// Add the template set to the cache, using the name of the page 
		// (like 'home.page.tmpl') as the key. 
		cache[name] = ts 
	}
	// Return the map. 
	return cache, nil 
}
