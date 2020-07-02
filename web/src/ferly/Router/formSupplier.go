package main

import (
	"fmt"
	"net/http"
)

func formData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Halaman Tidak Ditemukan", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GATE":
		http.ServeFile(w, r, "formSPL.html")

	case "POST":
		var kesalahan = r.ParseForm()
		if kesalahan != nil {
			fmt.Fprintln(w, "Ada Kesalahan :", kesalahan)
			return
		}

		var kdspl = r.FormValue("kdspl")
		var nmspl = r.FormValue("nmspl")
		var alamat = r.FormValue("alamat")
		var nohp = r.FormValue("nohp")
		fmt.Fprintln(w, "Kode Supplier = ", kdspl)
		fmt.Fprintln(w, "Nama Supplier = ", nmspl)
		fmt.Fprintln(w, "Alamat Supplier = ", alamat)
		fmt.Fprintln(w, "No HP = ", nohp)

	default:
		fmt.Fprintln(w, "Maaf Method yang didukung hanya GET dan POST.")
	}
}
func main() {
	http.HandleFunc("/", formData)

	fmt.Println("Web berjalan di alamat http:/localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
