package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("./assets/"))))
	http.ListenAndServe(":8080", nil)
}
func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index page",
	}
	err := tpl.ExecuteTemplate(w, "index.html", pd)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return

	}

	fmt.Println(req.URL.Path)
	fmt.Println("Success")
}
func about(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "About page",
	}
	err := tpl.ExecuteTemplate(w, "about.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(req.URL.Path)
	fmt.Println("Success")
}
func apply(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Index page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(req.URL.Path)
	fmt.Println("Success")
}
func contact(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Contact page",
	}
	err := tpl.ExecuteTemplate(w, "contact.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Println(req.URL.Path)
	fmt.Println("Success")
}
