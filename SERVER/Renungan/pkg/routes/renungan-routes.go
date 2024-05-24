package routes

import (
	"github.com/20WilliamPanjaitan/renungan/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi RegisterRenungansRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var RenunganRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi CreateRenungan dari package controllers.
	router.HandleFunc("/renungan", controllers.CreateRenungan).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi GetRenungan dari package controllers.
	router.HandleFunc("/renungan", controllers.GetRenungan).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetRenunganById dari package controllers.
	router.HandleFunc("/renungan/{id}", controllers.GetRenunganById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi UpdateRenungan dari package controllers.
	router.HandleFunc("/renungan/{id}", controllers.UpdateRenungan).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi DeleteRenungan dari package controllers.
	router.HandleFunc("/renungan/{id}", controllers.DeleteRenungan).Methods("DELETE")
}
