package main

import (
    "ascii-art-web/static"
    "html/template"
    "net/http"
    "bytes"
)

type banner struct {
    ASCIIArt string
}

type responseWriterWithBuffer struct {
    http.ResponseWriter
    buf *bytes.Buffer
}

func (rw *responseWriterWithBuffer) Write(p []byte) (int, error) {
    return rw.buf.Write(p)
}

func asciiart(input, fonts string) string {
    fontFolder := "static/fonts"
    var buf bytes.Buffer
    rw := &responseWriterWithBuffer{buf: &buf}
    ascii_art.AsciiArt(input, fonts, rw, fontFolder)
    return buf.String()
}

func handleRequest() {
    http.HandleFunc("/", mainPage)
    http.HandleFunc("/ascii-art", asciiArtHandler)
    http.ListenAndServe(":8080", nil)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
    tmpl, err := template.ParseFiles("templates/index.html", "templates/styles/styles.css")

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, "")
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html", "templates/styles/styles.css")

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
      
    r.ParseForm()
    input := r.FormValue("input")
    fonts := r.FormValue("fonts")

    if input == "" || fonts == "" {
        http.Error(w, "Missing input data", http.StatusBadRequest)
        return
    }

    asciiArtResult := asciiart(input, fonts)

    pageData := banner{
        ASCIIArt: asciiArtResult,
    }

    tmpl.Execute(w, pageData)
}

func main() {
    handleRequest()
}


