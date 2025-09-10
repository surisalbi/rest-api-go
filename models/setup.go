package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// membuat variabel global untuk database
var DB *gorm.DB

func ConnectDatabase() {
	// menghubungkan dengan database mysql
	db, err := gorm.Open(mysql.Open("root:lHvyuECorZazOiDQxygnblLCcVDLaURc@tcp(mysql.railway.internal:3306)/railway"))

	//cek koneksi ke database
	if err != nil {
		panic(err)
	}

	// migrasi tabel jika tabel di database kosong
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Member{})
	DB = db
}