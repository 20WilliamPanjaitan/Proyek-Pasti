package models

import (
"github.com/20WilliamPanjaitan/parhalado/pkg/config"
"github.com/jinzhu/gorm"

)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Parhalado.
type Parhalado struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
Nama string `json:"nama" gorm:"column:nama"`
TanggalLahir string `json:"tanggal_lahir" gorm:"column:tanggal_lahir"`
Jabatan string `json:"jabatan" gorm:"column:jabatan"`
Alamat string `json:"alamat" gorm:"column:alamat"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Parhalado{}) // Melakukan migrasi struktur data Parhalado ke dalam database.
}

// Fungsi untuk membuat Parhalado baru.
func (b *Parhalado) CreateParhalado() *Parhalado {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Parhalado baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Parhalados.
func GetAllParhalados() []Parhalado {
var Parhalados []Parhalado
db.Find(&Parhalados) // Mendapatkan semua data Parhalados dari database.
return Parhalados
}



func GetParhaladoById(id int64) (*Parhalado, *gorm.DB) {
var getParhalado Parhalado
db := db.Where("id = ?", id).Find(&getParhalado) // Mengambil data mahasiswa berdasarkan NIM.
return &getParhalado, db
}

// Fungsi untuk menghapus Parhalado berdasarkan kode Parhalado.
func DeleteParhaladoById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Parhalado{}).Error; err != nil {
return err
}
return nil
}