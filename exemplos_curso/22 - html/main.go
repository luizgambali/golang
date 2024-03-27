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

func Home(w http.ResponseWriter, r *http.Request) {

	u := usuario{"João", "joao.silva@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {

	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", Home)

	fmt.Println("Iniciando em http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
