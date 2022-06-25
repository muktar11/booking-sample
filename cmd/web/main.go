package main

import (
	"fmt"
	"booking/pkg/config"
	"booking/pkg/handlers"
	"booking/pkg/render"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)


const portNumber = ":8000"
var app config.AppConfig 
var session *scs.SessionManager



func main() { 
	var app config.AppConfig 

	//change  this to true when in production
	app.InProduction = false 
	
	session := scs.New()
	session.Lifetime = 24  * time.Hour
	session.Cookie.Persist = true 
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProductjion  

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("can not create template cache")
	}
	app.TemplaceCache = tc 
	app.UseCache = false 

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting  applocate on port %s", portNumber))

	// _ = http.Listen and Serve Port Number 

	srv := &http.Server {
		Addr : portNumber, 
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	
}