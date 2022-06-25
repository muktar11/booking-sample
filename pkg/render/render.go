package render

import (
	"bytes"
	"fmt"
	"goGoFirst/pkg/config"
	"goGoFirst/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig
//NewTemplates set the config for the template package 
func NewTemplates(a *config.AppConfig){
	app = a
}

//a funcion you use in template
var functions = template.FuncMap{
                                           
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td 
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td handlers, TemplateData) {

	var tc map[string]*template.Template 
	if app.UseCache {
		tc = app.TemplaceCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil{
		fmt.Println( "Error writing template to browser", err)
		
	}
}
//createTemplateCache creates a template cache as a map
//create
func CreateTemplateCache(w http.ResponseWriter, tmpl string) (map[string]*template.Template,error) {
	myCache := map[string]*template.Template{} 

	pages, err := filepath.Glob("./templates/*page.tmpl")
	
	if err != nil{
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache, err 
		}

		matches, err := filepath.Glob("./templates/*.layout,tmpl") 
		if err != nil{
			return myCache, err 
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.html")
			if err != nil{
				return myCache, err 
			}
		}
		myCache[name] = ts 
	}
	return myCache, nil 
}