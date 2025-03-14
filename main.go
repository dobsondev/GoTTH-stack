package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/dobsondev/gotth-stack/handlers"
	"github.com/dobsondev/gotth-stack/templ/pages"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	page := pages.DemoPage()
	http.Handle("/", templ.Handler(page))

	http.HandleFunc("/api/hello", handlers.HelloHandler)

	fmt.Println("App running on :3000")
	http.ListenAndServe(":3000", nil)
}
