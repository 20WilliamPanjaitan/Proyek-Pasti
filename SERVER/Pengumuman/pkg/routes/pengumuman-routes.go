package routes

import (
	"github.com/20WilliamPanjaitan/pengumuman/pkg/controllers"
	"github.com/gorilla/mux"
)

// Fungsi RegisterpengumumansRoutes mendaftarkan semua rute yang terkait dengan entitas mahasiswa.
// Fungsi ini menerima objek router dari Gorilla Mux untuk mendaftarkan rutenya.
var PengumumanRoutes = func(router *mux.Router) {
	// Menyambungkan rute HTTP POST untuk membuat mahasiswa baru ke fungsi Createpengumuman dari package controllers.
	router.HandleFunc("/pengumuman", controllers.CreatePengumuman).Methods("POST")
	// Menyambungkan rute HTTP GET untuk mendapatkan daftar semua mahasiswa ke fungsi Getpengumuman dari package controllers.
	router.HandleFunc("/pengumuman", controllers.GetPengumuman).Methods("GET")
	// Menyambungkan rute HTTP GET untuk mendapatkan detail mahasiswa berdasarkan ID ke fungsi GetpengumumanById dari package controllers.
	router.HandleFunc("/pengumuman/{id}", controllers.GetPengumumanById).Methods("GET")
	// Menyambungkan rute HTTP PUT untuk memperbarui informasi mahasiswa berdasarkan ID ke fungsi Updatepengumuman dari package controllers.
	router.HandleFunc("/pengumuman/{id}", controllers.UpdatePengumuman).Methods("PUT")
	// Menyambungkan rute HTTP DELETE untuk menghapus mahasiswa berdasarkan ID ke fungsi Deletepengumuman dari package controllers.
	router.HandleFunc("/pengumuman/{id}", controllers.DeletePengumuman).Methods("DELETE")
}
