package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/renungan/pkg/models"
	"github.com/20WilliamPanjaitan/renungan/pkg/utils"
	"github.com/gorilla/mux"
)

var NewRenungan models.Renungan

func GetRenungan(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllRenungans dari package models.
	newRenungans := models.GetAllRenungans() // Ambil hanya elemen pertama dari slice
	// Mengonversi Renungan menjadi format JSON.
	res, err := json.Marshal(newRenungans)
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

func GetRenunganById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	RenunganId := vars["id"]
	// Mengonversi ID Renungan menjadi tipe data int64.
	IDRenungan, err := strconv.ParseInt(RenunganId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetRenunganById dari package models untuk mendapatkan detail Renungan berdasarkan ID.
	renunganDetails, _ := models.GetRenunganById(IDRenungan)
	// Mengonversi detail Renungan menjadi format JSON.
	res, _ := json.Marshal(renunganDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreateRenungan(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Renungan baru.
	CreateRenungan := &models.Renungan{}
	// Parsing body request menjadi objek Renungan.
	utils.ParseBody(r, CreateRenungan)
	// Memanggil fungsi CreateRenungan dari package models untuk membuat Renungan baru.
	b := CreateRenungan.CreateRenungan()
	// Mengonversi hasil pembuatan Renungan baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeleteRenungan(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	RenunganId := vars["id"]
	// Mengonversi ID Renungan menjadi tipe data int64.
	IDRenungan, err := strconv.ParseInt(RenunganId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeleteRenungan dari package models untuk menghapus Renungan berdasarkan ID.
	Renungan := models.DeleteRenunganById(IDRenungan)
	// Mengonversi hasil penghapusan Renungan menjadi format JSON.
	res, _ := json.Marshal(Renungan)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdateRenungan(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updateRenungan baru.
	var updateRenungan = &models.Renungan{}
	// Parsing body request menjadi objek Renungan.
	utils.ParseBody(r, updateRenungan)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	RenunganId := vars["id"]
	// Mengonversi ID Renungan menjadi tipe data int64.
	IDRenungan, err := strconv.ParseInt(RenunganId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetRenunganById dari package models untuk mendapatkan detail Renungan berdasarkan ID.
	renunganDetails, db := models.GetRenunganById(IDRenungan)

	// Memperbarui informasi Renungan sesuai dengan data yang diberikan dalam body request.
	if updateRenungan.Judul != "" {
		renunganDetails.Judul = updateRenungan.Judul
	}
	if updateRenungan.Tanggal != "" {
		renunganDetails.Tanggal = updateRenungan.Tanggal
	}
	if updateRenungan.IsiRenungan != "" {
		renunganDetails.IsiRenungan = updateRenungan.IsiRenungan
	}
	if updateRenungan.Status != "" {
		renunganDetails.Status = updateRenungan.Status
	}
	if updateRenungan.AyatRenungan != "" {
		renunganDetails.AyatRenungan = updateRenungan.AyatRenungan
	}

	// Menyimpan perubahan informasi Renungan ke dalam database.
	db.Save(&renunganDetails)
	// Mengonversi detail Renungan yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(renunganDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
