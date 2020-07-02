package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

// membuat type mahasiswa dengan struktur
type mahasiswa struct {
	Nim    string
	Nama   string
	Progdi string
	Smt    int
}

type response struct {
	Status bool
	Pesan  string
	Data   []mahasiswa
}

// Membuat fungsi koneksi dengan sql
// sintax -> sql.open("mysql", "user : password@tep (host:port)/nama_database")
// karena bawaan xampp password kosong jadi dikosongkan aja

func koneksi() (*sql.DB, error) {
	db, salahe := sql.Open("mysql", "user:password@tcp(127.0.0.3:3306/cloud_udb")
	if salahe != nil {
		return nil, salahe
	}

	return db, nil
}

// fungsi Tampil Data

func tampil(pesane string) response {
	db, salahe := koneksi()
	if salahe != nil {

		return response{
			Status: false,
			Pesan:  "Gagal Koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	dataMhs, salahe := db.Query("Select * From mahasiswa")
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Koneksi : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()

	var hasil []mahasiswa

	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Progdi, &mhs.Smt)

		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal Baca : " + salahe.Error(),
				Data:   []mahasiswa{},
			}
		}

		hasil = append(hasil, mhs)
	}
	salahe = dataMhs.Err()

	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan : " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  pesane,
		Data:   hasil,
	}
}

func getMhs(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {

		return response{
			Status: false,
			Pesan:  "Gagal Koneksi :" + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	dataMhs, salahe := db.Query("select * from mahasiswa where nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()

	var hasil []mahasiswa

	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Progdi, &mhs.Smt)

		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal Baca" + salahe.Error(),
				Data:   []mahasiswa{},
			}
		}
		hasil = append(hasil, mhs)
	}
	salahe = dataMhs.Err()

	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Tampil",
		Data:   hasil,
	}
}

// Fungsi tambah Data

func tambah(nim string, nama string, progdi string, smt string) response {
	db, salahe := koneksi()
	if salahe != nil {

		return response{
			Status: false,
			Pesan:  "Gagal Koneksi:" + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	_, salahe = db.Exec("insert into mahasiswa values (?, ?, ?, ?)", nim, nama, progdi, smt)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Insert: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  "Berhasil Tambah",
		Data:   []mahasiswa{},
	}
}

// Fungsi Ubah data
func ubah(nim string, nama string, progdi string, smt string) response {
	db, salahe := koneksi()
	if salahe != nil {

		return response{
			Status: false,
			Pesan:  "Gagal Koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	_, salahe = db.Exec("update mahasiswa set nama =?, progdi =?, smt=? where nim=?", nama, progdi, smt, nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Update: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Ubah: ",
		Data:   []mahasiswa{},
	}
}

//Fungsi Hapus data

func hapus(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {

		return response{
			Status: false,
			Pesan:  "Gagal Koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	_, salahe = db.Exec("delete  from mahasiswa where nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Query Delete: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:  "Berhasil Hapus",
		Data:   []mahasiswa{},
	}
}

func kontroler(w http.ResponseWriter, r *http.Request) {

	var tampilHtml, salaheTampil = template.ParseFiles("template/tampil.html")
	if salaheTampil != nil {
		fmt.Println(salaheTampil.Error())
		return
	}

	var tambahHtml, salaheTambah = template.ParseFiles("template/tampil.html")
	if salaheTambah != nil {
		fmt.Println(salaheTambah.Error())
		return
	}
	var ubahHtml, salaheUbah = template.ParseFiles("template/tampil.html")
	if salaheUbah != nil {
		fmt.Println(salaheUbah.Error())
		return
	}
	var hapusHtml, salaheHapus = template.ParseFiles("template/tampil.html")
	if salaheHapus != nil {
		fmt.Println(salaheHapus.Error())
		return
	}

	switch r.Method {

	case "GET":

		aksi := r.URL.Query()["aksi"]
		if len(aksi) == 0 {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		} else if aksi[0] == "tambah" {
			tambahHtml.Execute(w, nil)
		} else if aksi[0] == "ubah" {
			nim := r.URL.Query()["nim"]
			ubahHtml.Execute(w, getMhs(nim[0]))
		} else if aksi[0] == "hapus" {
			nim := r.URL.Query()["nim"]
			hapusHtml.Execute(w, getMhs(nim[0]))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		}

	case "Post":
		var salahe = r.ParseForm()
		if salahe != nil {
			fmt.Fprintln(w, "Kesalahan: ", salahe)
			return
		}

		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		var progdi = r.FormValue("progdi")
		var smt = r.FormValue("smt")

		var aksi = r.URL.Path

		if aksi == "/tambah" {
			var hasil = tambah(nim, nama, progdi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/ubah" {
			var hasil = ubah(nim, nama, progdi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/hapus" {
			var hasil = hapus(nim)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else {
			tampilHtml.Execute(w, tampil("berhasil Tampil"))
		}
	default:
		fmt.Fprintln(w, "Maaf. Method yang didukung hanyak GET dan POST")
	}
}

func main() {

	http.HandleFunc("/", kontroler)

	fmt.Println("Server Berjalan Di Port http://localhost:8080/?aksi=tambah")
	http.ListenAndServe(":8080", nil)
}
