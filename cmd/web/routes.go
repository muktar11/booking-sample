package main

import (
	"booking/pkg/config"
	"booking/pkg/handlers"
	"net/http"

	"gihtub.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes (app *config.AppConfig)  http.Handler {
	mux := pat.New()
	
//	mux.Get("/", http.HandleFunc(handlers.Repo.Home))
//	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	//handker for crassh 
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux 
}