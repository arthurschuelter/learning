package main

import (
	"chi-api/internal/config"
	"chi-api/internal/handler"
	"chi-api/internal/repository"
	"chi-api/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	PrintBanner()
	cfg := config.LoadConfig()
	db, err := cfg.Connect()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// --- Links Module Setup ---
	linksRepo := repository.NewLinkRepo(db)
	linksService := service.NewLinkService(linksRepo)
	linksHandler := handler.NewLinksHandler(linksService)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from Chi v0"))
	})

	// --- Links Routes ---
	r.Route("/", func(r chi.Router) {
		r.Get("/", linksHandler.GetLinks)
		r.Get("/{code}", linksHandler.FindLink)
		r.Post("/", linksHandler.SaveLink)
	})

	fmt.Println("Chi v0 -- Listening on localhost:" + cfg.ApiPort)
	http.ListenAndServe(":"+cfg.ApiPort, r)
}

func PrintBanner() {
	banner := `
(Chi @ v0)
 ________  ________          ________  ________  ___     
|\   ____\|\   __  \        |\   __  \|\   __  \|\  \    
\ \  \___|\ \  \|\  \       \ \  \|\  \ \  \|\  \ \  \   
 \ \  \  __\ \  \\\  \       \ \   __  \ \   ____\ \  \  
  \ \  \|\  \ \  \\\  \       \ \  \ \  \ \  \___|\ \  \ 
   \ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\
    \|_______|\|_______|        \|__|\|__|\|__|     \|__|
	`

	fmt.Println(banner)
}
