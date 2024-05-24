package routes

import (
	"github.com/20WilliamPanjaitan/parhalado/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi RegisterParhaladosRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var ParhaladoRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi CreateParhalado dari package controllers.
	router.HandleFunc("/parhalado", controllers.CreateParhalado).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi GetParhalado dari package controllers.
	router.HandleFunc("/parhalado", controllers.GetParhalado).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetParhaladoById dari package controllers.
	router.HandleFunc("/parhalado/{id}", controllers.GetParhaladoById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi UpdateParhalado dari package controllers.
	router.HandleFunc("/parhalado/{id}", controllers.UpdateParhalado).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi DeleteParhalado dari package controllers.
	router.HandleFunc("/parhalado/{id}", controllers.DeleteParhalado).Methods("DELETE")
}
