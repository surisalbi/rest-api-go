package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/surisalbi/rest-api-go/controllers/authorcontroller"
	"github.com/surisalbi/rest-api-go/controllers/bookcontroller"
	"github.com/surisalbi/rest-api-go/controllers/membercontroller"
	"github.com/surisalbi/rest-api-go/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	
	app.Get("/api/books", bookcontroller.Index)
	app.Get("/api/books/:id", bookcontroller.Show)
	app.Post("/api/books/", bookcontroller.Create)
	app.Put("/api/books/:id", bookcontroller.Update)
	app.Delete("/api/books/:id", bookcontroller.Delete)

	app.Get("/api/authors", authorcontroller.Index)
	app.Get("/api/authors/:id", authorcontroller.Show)
	app.Post("/api/authors/", authorcontroller.Create)
	app.Put("/api/authors/:id", authorcontroller.Update)
	app.Delete("/api/authors/:id", authorcontroller.Delete)

	app.Get("/api/members", membercontroller.Index)
	app.Get("/api/members/:id", membercontroller.Show)
	app.Post("/api/members/", membercontroller.Create)
	app.Put("/api/members/:id", membercontroller.Update)
	app.Delete("/api/members/:id", membercontroller.Delete)

	app.Listen(":8000")
}