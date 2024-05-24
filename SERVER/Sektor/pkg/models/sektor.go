package models

import (
"github.com/20WilliamPanjaitan/sektor/pkg/config"
"github.com/jinzhu/gorm"

)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Sektor.
type Sektor struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
NamaSektor string `json:"nama_sektor" gorm:"column:nama_sektor"`
Deskripsi string `json:"deskripsi" gorm:"column:deskripsi"`
Alamat string `json:"alamat" gorm:"column:alamat"`
Kontak string `json:"kontak" gorm:"column:kontak"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Sektor{}) // Melakukan migrasi struktur data Sektor ke dalam database.
}

// Fungsi untuk membuat Sektor baru.
func (b *Sektor) CreateSektor() *Sektor {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Sektor baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Sektors.
func GetAllSektors() []Sektor {
var Sektors []Sektor
db.Find(&Sektors) // Mendapatkan semua data Sektors dari database.
return Sektors
}



func GetSektorById(id int64) (*Sektor, *gorm.DB) {
var getSektor Sektor
db := db.Where("id = ?", id).Find(&getSektor) // Mengambil data mahasiswa berdasarkan id.
return &getSektor, db
}

// Fungsi untuk menghapus Sektor berdasarkan kode Sektor.
func DeleteSektorById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Sektor{}).Error; err != nil {
return err
}
return nil
}