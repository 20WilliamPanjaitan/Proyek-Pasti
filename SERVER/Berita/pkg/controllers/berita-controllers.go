package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/berita/pkg/models"
	"github.com/20WilliamPanjaitan/berita/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBerita models.Berita

func GetBerita(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllBeritas dari package models.
	newBeritas := models.GetAllBeritas() // Ambil hanya elemen pertama dari slice
	// Mengonversi Berita menjadi format JSON.
	res, err := json.Marshal(newBeritas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func GetBeritaById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	BeritaId := vars["id"]
	// Mengonversi ID Berita menjadi tipe data int64.
	IDBerita, err := strconv.ParseInt(BeritaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetBeritaById dari package models untuk mendapatkan detail Berita berdasarkan ID.
	beritaDetails, _ := models.GetBeritaById(IDBerita)
	// Mengonversi detail Berita menjadi format JSON.
	res, _ := json.Marshal(beritaDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreateBerita(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Berita baru.
	CreateBerita := &models.Berita{}
	// Parsing body request menjadi objek Berita.
	utils.ParseBody(r, CreateBerita)
	// Memanggil fungsi CreateBerita dari package models untuk membuat Berita baru.
	b := CreateBerita.CreateBerita()
	// Mengonversi hasil pembuatan Berita baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeleteBerita(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	BeritaId := vars["id"]
	// Mengonversi ID Berita menjadi tipe data int64.
	IDBerita, err := strconv.ParseInt(BeritaId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeleteBerita dari package models untuk menghapus Berita berdasarkan ID.
	Berita := models.DeleteBeritaById(IDBerita)
	// Mengonversi hasil penghapusan Berita menjadi format JSON.
	res, _ := json.Marshal(Berita)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdateBerita(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updateBerita baru.
	var updateBerita = &models.Berita{}
	// Parsing body request menjadi objek Berita.
	utils.ParseBody(r, updateBerita)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	BeritaId := vars["id"]
	// Mengonversi ID Berita menjadi tipe data int64.
	IDBerita, err := strconv.ParseInt(BeritaId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetBeritaById dari package models untuk mendapatkan detail Berita berdasarkan ID.
	beritaDetails, db := models.GetBeritaById(IDBerita)

	// Memperbarui informasi Berita sesuai dengan data yang diberikan dalam body request.
	if updateBerita.Judul != "" {
		beritaDetails.Judul = updateBerita.Judul
	}
	if updateBerita.Tanggal != "" {
		beritaDetails.Tanggal = updateBerita.Tanggal
	}
	if updateBerita.Keterangan != "" {
		beritaDetails.Keterangan = updateBerita.Keterangan
	}
	if updateBerita.Status != "" {
		beritaDetails.Status = updateBerita.Status
	}

	// Menyimpan perubahan informasi Berita ke dalam database.
	db.Save(&beritaDetails)
	// Mengonversi detail Berita yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(beritaDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
