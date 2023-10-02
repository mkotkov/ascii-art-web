package main

import (
	"ascii-art-web/static"
	"html/template"
	"fmt"
	"net/http"
)

func asciiart(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    input := r.FormValue("input")
    fonts := r.FormValue("fonts")
    
    // Проверка наличия данных в форме
    if input == "" || fonts == "" {
        http.Error(w, "Missing input data", http.StatusBadRequest)
        return
    }
    
    fontFolder := "static/fonts"
    ascii_art.AsciiArt(input, fonts, w, fontFolder)
}

func handleRequest() {
    http.HandleFunc("/", main_page)
    http.HandleFunc("/ascii-art", asciiart)
    http.ListenAndServe(":3535", nil)
}

func main_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("temlates/index.html")
	tmpl.Execute(w, nil)
}

func main() {
    handleRequest()
}