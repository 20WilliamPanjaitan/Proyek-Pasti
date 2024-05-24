package models

import (
"github.com/20WilliamPanjaitan/pengumuman/pkg/config"
"github.com/jinzhu/gorm"

)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Pengumuman.
type Pengumuman struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
Judul string `json:"judul" gorm:"column:judul"`
Tanggal string `json:"tanggal" gorm:"column:tanggal"`
Keterangan string `json:"keterangan" gorm:"column:keterangan"`
Status string `json:"status" gorm:"column:status"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Pengumuman{}) // Melakukan migrasi struktur data Pengumuman ke dalam database.
}

// Fungsi untuk membuat Pengumuman baru.
func (b *Pengumuman) CreatePengumuman() *Pengumuman {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Pengumuman baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Pengumumans.
func GetAllPengumumans() []Pengumuman {
var Pengumumans []Pengumuman
db.Find(&Pengumumans) // Mendapatkan semua data Pengumumans dari database.
return Pengumumans
}



func GetPengumumanById(id int64) (*Pengumuman, *gorm.DB) {
var getPengumuman Pengumuman
db := db.Where("id = ?", id).Find(&getPengumuman) // Mengambil data mahasiswa berdasarkan NIM.
return &getPengumuman, db
}

// Fungsi untuk menghapus Pengumuman berdasarkan kode Pengumuman.
func DeletePengumumanById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Pengumuman{}).Error; err != nil {
return err
}
return nil
}