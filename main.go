package main

import (
	"fmt"
	"net/http"

	"github.com/dobsondev/gotth-stack/handlers"
	"github.com/dobsondev/gotth-stack/templ/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	page := pages.DemoPage()
	router.Handle("/", templ.Handler(page))

	router.HandleFunc("/api/hello", handlers.HelloHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("App running on :3000")
	server.ListenAndServe()
}
