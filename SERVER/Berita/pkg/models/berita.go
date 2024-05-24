package models

import (
"github.com/20WilliamPanjaitan/berita/pkg/config"
"github.com/jinzhu/gorm"


)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Berita.
type Berita struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
Judul string `json:"judul" gorm:"column:judul"`
Keterangan string `json:"keterangan" gorm:"column:keterangan"`
Status string `json:"status" gorm:"column:status"`
Tanggal string `json:"tanggal" gorm:"column:tanggal"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Berita{}) // Melakukan migrasi struktur data Berita ke dalam database.
}

// Fungsi untuk membuat Berita baru.
func (b *Berita) CreateBerita() *Berita {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Berita baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Beritas.
func GetAllBeritas() []Berita {
var Beritas []Berita
db.Find(&Beritas) // Mendapatkan semua data Beritas dari database.
return Beritas
}



func GetBeritaById(id int64) (*Berita, *gorm.DB) {
var getBerita Berita
db := db.Where("id = ?", id).Find(&getBerita) // Mengambil data mahasiswa berdasarkan NIM.
return &getBerita, db
}

// Fungsi untuk menghapus Berita berdasarkan kode Berita.
func DeleteBeritaById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Berita{}).Error; err != nil {
return err
}
return nil
}