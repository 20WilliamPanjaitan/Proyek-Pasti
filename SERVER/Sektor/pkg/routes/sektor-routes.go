package routes

import (
	"github.com/20WilliamPanjaitan/sektor/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi RegisterSektorsRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var SektorRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi CreateSektor dari package controllers.
	router.HandleFunc("/sektor", controllers.CreateSektor).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi GetSektor dari package controllers.
	router.HandleFunc("/sektor", controllers.GetSektor).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetSektorById dari package controllers.
	router.HandleFunc("/sektor/{id}", controllers.GetSektorById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi UpdateSektor dari package controllers.
	router.HandleFunc("/sektor/{id}", controllers.UpdateSektor).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi DeleteSektor dari package controllers.
	router.HandleFunc("/sektor/{id}", controllers.DeleteSektor).Methods("DELETE")
}
