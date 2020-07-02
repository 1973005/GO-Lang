package main

//Membuat halaman route dan profil
import (
	"fmt"
	"net/http"
)

func profil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halaman Utama")
}

func main() {
	//Membuat halaman route dan profil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halaman Utama")
	})
	// http.HandleFunc("/profil", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Halaman Profil")
	// })
	http.HandleFunc("/profil", profil)
	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
