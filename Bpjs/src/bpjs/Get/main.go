package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintln(w, "Posting DiWeb! r.PostFrom = %v\n", r.PostForm)
		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		fmt.Fprintf(w, "Nama = %s\n", nama)
		fmt.Fprintf(w, "Alamat = %s\n", alamat)
	default:
		fmt.Fprintf(w, "Maaf Method yang didukung hanya GET dan POST.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Web berjalan di alamat http:/localhost:8080/\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	http.HandleFunc("/", formData)

// 	fmt.Println("Web berjalan di alamat http:/localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }
