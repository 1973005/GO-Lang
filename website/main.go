package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("./assets/"))))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}
