package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/sektor/pkg/models"
	"github.com/20WilliamPanjaitan/sektor/pkg/utils"
	"github.com/gorilla/mux"
)

var NewSektor models.Sektor

func GetSektor(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllSektors dari package models.
	newSektors := models.GetAllSektors() // Ambil hanya elemen pertama dari slice
	// Mengonversi Sektor menjadi format JSON.
	res, err := json.Marshal(newSektors)
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

func GetSektorById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	SektorId := vars["id"]
	// Mengonversi ID Sektor menjadi tipe data int64.
	IDSektor, err := strconv.ParseInt(SektorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetSektorById dari package models untuk mendapatkan detail Sektor berdasarkan ID.
	sektorDetails, _ := models.GetSektorById(IDSektor)
	// Mengonversi detail Sektor menjadi format JSON.
	res, _ := json.Marshal(sektorDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreateSektor(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Sektor baru.
	CreateSektor := &models.Sektor{}
	// Parsing body request menjadi objek Sektor.
	utils.ParseBody(r, CreateSektor)
	// Memanggil fungsi CreateSektor dari package models untuk membuat Sektor baru.
	b := CreateSektor.CreateSektor()
	// Mengonversi hasil pembuatan Sektor baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeleteSektor(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	SektorId := vars["id"]
	// Mengonversi ID Sektor menjadi tipe data int64.
	IDSektor, err := strconv.ParseInt(SektorId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeleteSektor dari package models untuk menghapus Sektor berdasarkan ID.
	Sektor := models.DeleteSektorById(IDSektor)
	// Mengonversi hasil penghapusan Sektor menjadi format JSON.
	res, _ := json.Marshal(Sektor)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdateSektor(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updateSektor baru.
	var updateSektor = &models.Sektor{}
	// Parsing body request menjadi objek Sektor.
	utils.ParseBody(r, updateSektor)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	SektorId := vars["id"]
	// Mengonversi ID Sektor menjadi tipe data int64.
	IDSektor, err := strconv.ParseInt(SektorId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetSektorById dari package models untuk mendapatkan detail Sektor berdasarkan ID.
	sektorDetails, db := models.GetSektorById(IDSektor)

	// Memperbarui informasi Sektor sesuai dengan data yang diberikan dalam body request.
	if updateSektor.NamaSektor != "" {
		sektorDetails.NamaSektor = updateSektor.NamaSektor
	}
	if updateSektor.Deskripsi != "" {
		sektorDetails.Deskripsi = updateSektor.Deskripsi
	}
	if updateSektor.Alamat != "" {
		sektorDetails.Alamat = updateSektor.Alamat
	}
	if updateSektor.Kontak != "" {
		sektorDetails.Kontak = updateSektor.Kontak
	}

	// Menyimpan perubahan informasi Sektor ke dalam database.
	db.Save(&sektorDetails)
	// Mengonversi detail Sektor yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(sektorDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
