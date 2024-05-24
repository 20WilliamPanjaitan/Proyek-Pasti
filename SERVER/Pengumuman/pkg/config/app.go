package config

import (
	"github.com/jinzhu/gorm"                  // Mengimpor paket GORM untuk ORM
	_"github.com/jinzhu/gorm/dialects/mysql" // Mengimpor driver MySQL untuk GORM
)

var (
	db *gorm.DB // Mendeklarasikan variabel tingkat paket untuk menyimpan koneksi database
)

// Connect membuat koneksi ke database MySQL.
func Connect() {
	// Mencoba membuka koneksi ke database MySQL menggunakan GORM
	d, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/hkbp_parparean_pengumuman?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // Jika koneksi gagal, panic dan menghentikan program
	}
	db = d // Menetapkan koneksi database ke variabel tingkat paket
}

// GetDB mengembalikan instance dari database yang terhubung.
func GetDB() *gorm.DB {
	return db // Mengembalikan instance koneksi database
}
