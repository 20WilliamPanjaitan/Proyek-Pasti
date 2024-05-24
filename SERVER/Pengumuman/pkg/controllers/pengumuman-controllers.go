package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/pengumuman/pkg/models"
	"github.com/20WilliamPanjaitan/pengumuman/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPengumuman models.Pengumuman

func GetPengumuman(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllPengumumans dari package models.
	newPengumumans := models.GetAllPengumumans() // Ambil hanya elemen pertama dari slice
	// Mengonversi Pengumuman menjadi format JSON.
	res, err := json.Marshal(newPengumumans)
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

func GetPengumumanById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	PengumumanId := vars["id"]
	// Mengonversi ID Pengumuman menjadi tipe data int64.
	IDPengumuman, err := strconv.ParseInt(PengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetPengumumanById dari package models untuk mendapatkan detail Pengumuman berdasarkan ID.
	pengumumanDetails, _ := models.GetPengumumanById(IDPengumuman)
	// Mengonversi detail Pengumuman menjadi format JSON.
	res, _ := json.Marshal(pengumumanDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreatePengumuman(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Pengumuman baru.
	CreatePengumuman := &models.Pengumuman{}
	// Parsing body request menjadi objek Pengumuman.
	utils.ParseBody(r, CreatePengumuman)
	// Memanggil fungsi CreatePengumuman dari package models untuk membuat Pengumuman baru.
	b := CreatePengumuman.CreatePengumuman()
	// Mengonversi hasil pembuatan Pengumuman baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeletePengumuman(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	PengumumanId := vars["id"]
	// Mengonversi ID Pengumuman menjadi tipe data int64.
	IDPengumuman, err := strconv.ParseInt(PengumumanId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeletePengumuman dari package models untuk menghapus Pengumuman berdasarkan ID.
	Pengumuman := models.DeletePengumumanById(IDPengumuman)
	// Mengonversi hasil penghapusan Pengumuman menjadi format JSON.
	res, _ := json.Marshal(Pengumuman)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdatePengumuman(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updatePengumuman baru.
	var updatePengumuman = &models.Pengumuman{}
	// Parsing body request menjadi objek Pengumuman.
	utils.ParseBody(r, updatePengumuman)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	PengumumanId := vars["id"]
	// Mengonversi ID Pengumuman menjadi tipe data int64.
	IDPengumuman, err := strconv.ParseInt(PengumumanId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetPengumumanById dari package models untuk mendapatkan detail Pengumuman berdasarkan ID.
	pengumumanDetails, db := models.GetPengumumanById(IDPengumuman)

	// Memperbarui informasi Pengumuman sesuai dengan data yang diberikan dalam body request.
	if updatePengumuman.Judul != "" {
		pengumumanDetails.Judul = updatePengumuman.Judul
	}
	if updatePengumuman.Tanggal != "" {
		pengumumanDetails.Tanggal = updatePengumuman.Tanggal
	}
	if updatePengumuman.Keterangan != "" {
		pengumumanDetails.Keterangan = updatePengumuman.Keterangan
	}
	if updatePengumuman.Status != "" {
		pengumumanDetails.Status = updatePengumuman.Status
	}

	// Menyimpan perubahan informasi Pengumuman ke dalam database.
	db.Save(&pengumumanDetails)
	// Mengonversi detail Pengumuman yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(pengumumanDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
