package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/calculate/", h2)

	log.Println("App starting...")
	log.Fatal(http.ListenAndServe(":8008", nil))
}

func calculateBMR() int {
	return 0
}
