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

type RouteType int

const (
	TemplRoute   RouteType = iota // 0
	HandlerRoute                  // 1
)

type Route struct {
	Name       string
	Path       string
	Type       RouteType
	TestStatus int
	Handler    interface{}
}

var routes = []Route{
	{
		Name:       "Home",
		Path:       "/",
		Type:       TemplRoute,
		TestStatus: http.StatusOK,
		Handler:    pages.DemoPage,
	},
	{
		Name:       "API - Hello",
		Path:       "/api/hello",
		Type:       HandlerRoute,
		TestStatus: http.StatusOK,
		Handler:    handlers.HelloHandler,
	},
}

func addRoutes(router *chi.Mux) {
	for _, route := range routes {
		switch route.Type {
		case TemplRoute:
			// For templ components
			templFn, ok := route.Handler.(func() templ.Component)
			if ok {
				router.Handle(route.Path, templ.Handler(templFn()))
			}
		case HandlerRoute:
			// For standard http handlers
			handlerFn, ok := route.Handler.(func(http.ResponseWriter, *http.Request))
			if ok {
				router.HandleFunc(route.Path, handlerFn)
			}
		}
	}
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	addRoutes(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("App running on :3000")
	server.ListenAndServe()
}
