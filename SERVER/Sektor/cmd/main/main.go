package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Package gorilla/mux digunakan untuk routing HTTP yang lebih fleksibel.

	// "github.com/jinzhu/gorm/dialects/mysql" // Package jinzhu/gorm/dialects/mysql digunakan untuk berinteraksi dengan database MySQL.
	"github.com/20WilliamPanjaitan/sektor/pkg/routes" // Package routes berisi definisi rute-rute aplikasi.
)

func main() {
	r := mux.NewRouter() // Membuat router baru menggunakan package Gorilla Mux.

	routes.SektorRoutes(r) // Mendaftarkan rute-rute untuk manajemen data mahasiswa.

	http.Handle("/", r) // Menetapkan router sebagai handler untuk semua permintaan pada path root ("/").

	fmt.Print("Starting Server localhost:9013") // Menampilkan pesan bahwa server dimulai di localhost:9013.

	log.Fatal(http.ListenAndServe("localhost:9013", r)) // Menggunakan http.ListenAndServe untuk menjalankan server HTTP di localhost:9013.
	// log.Fatal akan menampilkan pesan log dan keluar dari program jika terjadi kesalahan saat menjalankan server.
}
