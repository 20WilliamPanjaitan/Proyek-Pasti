package models

import (
"github.com/20WilliamPanjaitan/acara/pkg/config"
"github.com/jinzhu/gorm"

)

// Deklarasi variabel global untuk objek database.
var db *gorm.DB

// Definisi struktur data Acara.
type Acara struct {
gorm.Model // Struktur data yang diperlukan oleh GORM untuk model.
NamaAcara string `json:"nama_acara" gorm:"column:nama_acara"`
LokasiAcara string `json:"lokasi_acara" gorm:"column:lokasi_acara"`
JenisAcara string `json:"jenis_acara" gorm:"column:jenis_acara"`
WaktuPelaksanaan string `json:"waktu_pelaksanaan" gorm:"column:waktu_pelaksanaan"`
Keterangan string `json:"keterangan" gorm:"column:keterangan"`
}

// Fungsi init akan dipanggil saat package models di-import.
// Fungsi ini digunakan untuk melakukan koneksi ke database dan melakukan migrasi struktur data.
func init() {
config.Connect() // Menginisialisasi koneksi ke database.
db = config.GetDB() // Mendapatkan objek database yang telah terkoneksi.
db.AutoMigrate(&Acara{}) // Melakukan migrasi struktur data Acara ke dalam database.
}

// Fungsi untuk membuat Acara baru.
func (b *Acara) CreateAcara() *Acara {
db.NewRecord(b) // Memulai pencatatan transaksi baru.
db.Create(&b) // Membuat data Acara baru di database.
return b
}

// Fungsi untuk mendapatkan daftar semua Acaras.
func GetAllAcaras() []Acara {
var Acaras []Acara
db.Find(&Acaras) // Mendapatkan semua data Acaras dari database.
return Acaras
}



func GetAcaraById(id int64) (*Acara, *gorm.DB) {
var getAcara Acara
db := db.Where("id = ?", id).Find(&getAcara) // Mengambil data mahasiswa berdasarkan id.
return &getAcara, db
}

// Fungsi untuk menghapus Acara berdasarkan kode Acara.
func DeleteAcaraById(id int64) error {
// Mengatur nilai deleted_at untuk menandai soft-delete
if err := db.Where("id = ?", id).Delete(&Acara{}).Error; err != nil {
return err
}
return nil
}