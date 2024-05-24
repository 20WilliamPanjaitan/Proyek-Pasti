package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/parhalado/pkg/models"
	"github.com/20WilliamPanjaitan/parhalado/pkg/utils"
	"github.com/gorilla/mux"
)

var NewParhalado models.Parhalado

func GetParhalado(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllParhalados dari package models.
	newParhalados := models.GetAllParhalados() // Ambil hanya elemen pertama dari slice
	// Mengonversi Parhalado menjadi format JSON.
	res, err := json.Marshal(newParhalados)
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

func GetParhaladoById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	ParhaladoId := vars["id"]
	// Mengonversi ID Parhalado menjadi tipe data int64.
	IDParhalado, err := strconv.ParseInt(ParhaladoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetParhaladoById dari package models untuk mendapatkan detail Parhalado berdasarkan ID.
	parhaladoDetails, _ := models.GetParhaladoById(IDParhalado)
	// Mengonversi detail Parhalado menjadi format JSON.
	res, _ := json.Marshal(parhaladoDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreateParhalado(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Parhalado baru.
	CreateParhalado := &models.Parhalado{}
	// Parsing body request menjadi objek Parhalado.
	utils.ParseBody(r, CreateParhalado)
	// Memanggil fungsi CreateParhalado dari package models untuk membuat Parhalado baru.
	b := CreateParhalado.CreateParhalado()
	// Mengonversi hasil pembuatan Parhalado baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeleteParhalado(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	ParhaladoId := vars["id"]
	// Mengonversi ID Parhalado menjadi tipe data int64.
	IDParhalado, err := strconv.ParseInt(ParhaladoId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeleteParhalado dari package models untuk menghapus Parhalado berdasarkan ID.
	Parhalado := models.DeleteParhaladoById(IDParhalado)
	// Mengonversi hasil penghapusan Parhalado menjadi format JSON.
	res, _ := json.Marshal(Parhalado)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdateParhalado(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updateParhalado baru.
	var updateParhalado = &models.Parhalado{}
	// Parsing body request menjadi objek Parhalado.
	utils.ParseBody(r, updateParhalado)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	ParhaladoId := vars["id"]
	// Mengonversi ID Parhalado menjadi tipe data int64.
	IDParhalado, err := strconv.ParseInt(ParhaladoId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetParhaladoById dari package models untuk mendapatkan detail Parhalado berdasarkan ID.
	parhaladoDetails, db := models.GetParhaladoById(IDParhalado)

	// Memperbarui informasi Parhalado sesuai dengan data yang diberikan dalam body request.
	if updateParhalado.Nama != "" {
		parhaladoDetails.Nama = updateParhalado.Nama
	}
	if updateParhalado.TanggalLahir != "" {
		parhaladoDetails.TanggalLahir = updateParhalado.TanggalLahir
	}
	if updateParhalado.Jabatan != "" {
		parhaladoDetails.Jabatan = updateParhalado.Jabatan
	}
	if updateParhalado.Alamat != "" {
		parhaladoDetails.Alamat = updateParhalado.Alamat
	}

	// Menyimpan perubahan informasi Parhalado ke dalam database.
	db.Save(&parhaladoDetails)
	// Mengonversi detail Parhalado yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(parhaladoDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
