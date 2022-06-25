package handlers

import (
	"goGoFirst/pkg/config"
	"goGoFirst/pkg/render"
	"net/http"
)

//TemplateData holds data send from hanslers to templates
/*
type TemplateData struct {
	StringMap map[string]string  
	INtMap    map[string]int 
	FloatMap  map[string]float32
	Data     map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Error       string   
}
*/
// Repository the  repository used by the handlers 
var Repo *Repository 

//Repository is the repositiry type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a config.AppConfig) *Repository {
	return &Repository{
		App:a, 
	}
}
func Home(w http.ResponseWriter, r *http.Request){
	remoteIP := e.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &TemplateData{})
}

func About(w http.ResponseWriter, r *http.Request) {
	//perform some logic 
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again."
	
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP 
	
	//send the data dto the template 
	render.RenderTemplate(w, "about.page.html", &TemplateData{
		StringMap: stringMap,
	})
}
