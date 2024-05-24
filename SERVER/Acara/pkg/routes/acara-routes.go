package routes

import (
	"github.com/20WilliamPanjaitan/acara/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi Register acarasRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var AcaraRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi CreateAcara dari package controllers.
	router.HandleFunc("/acara", controllers.CreateAcara).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi GetAcara dari package controllers.
	router.HandleFunc("/acara", controllers.GetAcara).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetAcaraById dari package controllers.
	router.HandleFunc("/acara/{id}", controllers.GetAcaraById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi UpdateAcara dari package controllers.
	router.HandleFunc("/acara/{id}", controllers.UpdateAcara).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi DeleteAcara dari package controllers.
	router.HandleFunc("/acara/{id}", controllers.DeleteAcara).Methods("DELETE")
}
