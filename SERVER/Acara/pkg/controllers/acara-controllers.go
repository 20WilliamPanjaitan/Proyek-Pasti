package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/20WilliamPanjaitan/acara/pkg/models"
	"github.com/20WilliamPanjaitan/acara/pkg/utils"
	"github.com/gorilla/mux"
)

var NewAcara models.Acara

func GetAcara(w http.ResponseWriter, r *http.Request) {
	// Memanggil fungsi GetAllAcaras dari package models.
	newAcaras := models.GetAllAcaras() // Ambil hanya elemen pertama dari slice
	// Mengonversi Acara menjadi format JSON.
	res, err := json.Marshal(newAcaras)
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

func GetAcaraById(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	AcaraId := vars["id"]
	// Mengonversi ID Acara menjadi tipe data int64.
	IDAcara, err := strconv.ParseInt(AcaraId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetAcaraById dari package models untuk mendapatkan detail Acara berdasarkan ID.
	acaraDetails, _ := models.GetAcaraById(IDAcara)
	// Mengonversi detail Acara menjadi format JSON.
	res, _ := json.Marshal(acaraDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func CreateAcara(w http.ResponseWriter, r *http.Request) {
	// Membuat objek Acara baru.
	CreateAcara := &models.Acara{}
	// Parsing body request menjadi objek Acara.
	utils.ParseBody(r, CreateAcara)
	// Memanggil fungsi CreateAcara dari package models untuk membuat Acara baru.
	b := CreateAcara.CreateAcara()
	// Mengonversi hasil pembuatan Acara baru menjadi format JSON.
	res, _ := json.Marshal(b)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func DeleteAcara(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	AcaraId := vars["id"]
	// Mengonversi ID Acara menjadi tipe data int64.
	IDAcara, err := strconv.ParseInt(AcaraId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi DeleteAcara dari package models untuk menghapus Acara berdasarkan ID.
	Acara := models.DeleteAcaraById(IDAcara)
	// Mengonversi hasil penghapusan Acara menjadi format JSON.
	res, _ := json.Marshal(Acara)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}

func UpdateAcara(w http.ResponseWriter, r *http.Request) {
	// Membuat objek updateAcara baru.
	var updateAcara = &models.Acara{}
	// Parsing body request menjadi objek Acara.
	utils.ParseBody(r, updateAcara)
	// Mendapatkan variabel dari path URL menggunakan Gorilla Mux.
	vars := mux.Vars(r)
	AcaraId := vars["id"]
	// Mengonversi ID Acara menjadi tipe data int64.
	IDAcara, err := strconv.ParseInt(AcaraId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing") // Menampilkan pesan kesalahan jika terjadi kesalahan parsing.
	}
	// Memanggil fungsi GetAcaraById dari package models untuk mendapatkan detail Acara berdasarkan ID.
	acaraDetails, db := models.GetAcaraById(IDAcara)

	// Memperbarui informasi Acara sesuai dengan data yang diberikan dalam body request.
	if updateAcara.NamaAcara != "" {
		acaraDetails.NamaAcara = updateAcara.NamaAcara
	}
	if updateAcara.LokasiAcara != "" {
		acaraDetails.LokasiAcara = updateAcara.LokasiAcara
	}
	if updateAcara.JenisAcara != "" {
		acaraDetails.JenisAcara = updateAcara.JenisAcara
	}
	if updateAcara.WaktuPelaksanaan != "" {
		acaraDetails.WaktuPelaksanaan = updateAcara.WaktuPelaksanaan
	}
	if updateAcara.Keterangan != "" {
		acaraDetails.Keterangan = updateAcara.Keterangan
	}

	// Menyimpan perubahan informasi Acara ke dalam database.
	db.Save(&acaraDetails)
	// Mengonversi detail acara yang telah diperbarui menjadi format JSON.
	res, _ := json.Marshal(acaraDetails)
	// Menetapkan header Content-Type sebagai application/json.
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan status kode OK (200) ke client.
	w.WriteHeader(http.StatusOK)
	// Mengirimkan respon JSON ke client.
	w.Write(res)
}
