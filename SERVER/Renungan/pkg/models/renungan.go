package models

import (
"github.com/20WilliamPanjaitan/renungan/pkg/config"
"github.com/jinzhu/gorm"

)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Renungan.
type Renungan struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
Judul string `json:"judul" gorm:"column:judul"`
IsiRenungan string `json:"isi_renungan" gorm:"column:isi_renungan"`
AyatRenungan string `json:"ayat_renungan" gorm:"column:ayat_renungan"`
Status string `json:"status" gorm:"column:status"`
Tanggal string `json:"tanggal" gorm:"column:tanggal"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Renungan{}) // Melakukan migrasi struktur data Renungan ke dalam database.
}

// Fungsi untuk membuat Renungan baru.
func (b *Renungan) CreateRenungan() *Renungan {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Renungan baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Renungans.
func GetAllRenungans() []Renungan {
var Renungans []Renungan
db.Find(&Renungans) // Mendapatkan semua data Renungans dari database.
return Renungans
}



func GetRenunganById(id int64) (*Renungan, *gorm.DB) {
var getRenungan Renungan
db := db.Where("id = ?", id).Find(&getRenungan) // Mengambil data mahasiswa berdasarkan id.
return &getRenungan, db
}

// Fungsi untuk menghapus Renungan berdasarkan kode Renungan.
func DeleteRenunganById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Renungan{}).Error; err != nil {
return err
}
return nil
}