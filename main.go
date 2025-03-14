package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	templates "github.com/dobsondev/gotth-stack/templ"
)

func main() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles"))))

	component := templates.Hello("Alex")
	http.Handle("/", templ.Handler(component))

	http.Handle("/buttonClicked", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Button clicked")
		fmt.Fprintf(w, "Button clicked!")
	}))

	fmt.Println("App running on :3000")
	http.ListenAndServe(":3000", nil)
}
