package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/surisalbi/rest-api-go/controllers/authorcontroller"
	"github.com/surisalbi/rest-api-go/controllers/bookcontroller"
	"github.com/surisalbi/rest-api-go/controllers/membercontroller"
	"github.com/surisalbi/rest-api-go/models"
)

func main() {
	// koneksi database (di dalam models.ConnectDatabase kamu bisa ambil MYSQL_URL)
	models.ConnectDatabase()

	app := fiber.New()

	// pakai env PORT dari Railway, fallback ke 3000 untuk local
	port := os.Getenv("3306")
	if port == "" {
		port = "3000"
	}

	// middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// routes books
	app.Get("/api/books", bookcontroller.Index)
	app.Get("/api/books/:id", bookcontroller.Show)
	app.Post("/api/books", bookcontroller.Create)
	app.Put("/api/books/:id", bookcontroller.Update)
	app.Delete("/api/books/:id", bookcontroller.Delete)

	// routes authors
	app.Get("/api/authors", authorcontroller.Index)
	app.Get("/api/authors/:id", authorcontroller.Show)
	app.Post("/api/authors", authorcontroller.Create)
	app.Put("/api/authors/:id", authorcontroller.Update)
	app.Delete("/api/authors/:id", authorcontroller.Delete)

	// routes members
	app.Get("/api/members", membercontroller.Index)
	app.Get("/api/members/:id", membercontroller.Show)
	app.Post("/api/members", membercontroller.Create)
	app.Put("/api/members/:id", membercontroller.Update)
	app.Delete("/api/members/:id", membercontroller.Delete)

	// listen ke port Railway
	log.Printf("Server running on port %s 🚀", port)
	log.Fatal(app.Listen(":" + port))
}
