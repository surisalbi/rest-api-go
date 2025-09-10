package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// membuat variabel global untuk database
var DB *gorm.DB

func ConnectDatabase() {
	// menghubungkan dengan database mysql
	dsn := os.Getenv("mysql://root:cfQJevUGGDUfPbvfXhImdnTvCbiwdHTd@mysql.railway.internal:3306/railway") // Railway kasih env ini
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Gagal koneksi database!")
    }

	// migrasi tabel jika tabel di database kosong
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Member{})
	DB = db
}