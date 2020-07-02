package main

import (
	"fmt"
	"log"
	"net/http"
)

func FormData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Halaman Tidak Dapat Ditemukan.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET"
		http.ServeFile(w, r, "template.html")

	case "POST"
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
		}
		var Nik = r.FormValue("nik")
		var Name = r.FormValue("nama")
		var Tempat = r.FormValue("ttl")
		var Tanggal = r.FormValue("tgl")
		var Alamat = r.FormValue("alamat")
		var RTRW = r.FormValue("rtrw")
		var Kelurahan = r.FormValue("kelurahan")
		var Kecamatan = r.FormValue("kecamatan")
		var Agama = r.FormValue("agama")
		var Pekerjaan = r.FormValue("kerja")
		var Kewarganegaraan = r.FormValue("negara")
		var Golongan = r.FormValue("golongan")

		fmt.Fprintln(w, "nik = ", NIK)
		fmt.Fprintln(w, "nama")
	}
}
