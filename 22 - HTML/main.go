package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := usuario{"Alfredo", "alfredneto@hotmail.com"}
		templates.ExecuteTemplate(w, "home.html", user)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
