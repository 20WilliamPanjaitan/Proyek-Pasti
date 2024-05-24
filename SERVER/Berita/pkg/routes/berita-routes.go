package routes

import (
	"github.com/20WilliamPanjaitan/berita/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi RegisterBeritasRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var BeritaRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi CreateBerita dari package controllers.
	router.HandleFunc("/berita", controllers.CreateBerita).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi GetBerita dari package controllers.
	router.HandleFunc("/berita", controllers.GetBerita).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetBeritaById dari package controllers.
	router.HandleFunc("/berita/{id}", controllers.GetBeritaById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi UpdateBerita dari package controllers.
	router.HandleFunc("/berita/{id}", controllers.UpdateBerita).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi DeleteBerita dari package controllers.
	router.HandleFunc("/berita/{id}", controllers.DeleteBerita).Methods("DELETE")
}
