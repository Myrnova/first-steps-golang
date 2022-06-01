package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Name  string
	Email string
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"João",
		"joão.pedro@gmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", u) //terceiro parametro é os dados que podemos passar para o arquivo, no nosso caso é estatico
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
