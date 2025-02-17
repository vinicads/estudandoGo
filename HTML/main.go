package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	fmt.Println("Renderizando HTML com HTTP em GO")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", nil)
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}