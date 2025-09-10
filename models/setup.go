package models

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil dari environment Railway
	mysqlURL := os.Getenv("MYSQL_URL")

	var dsn string
	if mysqlURL != "" {
		// parse mysql://user:pass@host:port/dbname
		parsedURL, err := url.Parse(mysqlURL)
		if err != nil {
			log.Fatal("Invalid MYSQL_URL:", err)
		}

		user := parsedURL.User.Username()
		pass, _ := parsedURL.User.Password()
		host := parsedURL.Hostname()
		port := parsedURL.Port()
		dbname := parsedURL.Path[1:] // buang '/' di depan

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, dbname)
	} else {
		// fallback ke lokal
		dsn = "root:root@tcp(localhost:8889)/rest_api_go?charset=utf8mb4&parseTime=True&loc=Local"
		log.Println("MYSQL_URL not found, using local DB")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// auto migrate
	db.AutoMigrate(&Book{}, &Author{}, &Member{})
	DB = db

	log.Println("Database connected 🚀")
}
