package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thanh/go-book1/database"
	"github.com/thanh/go-book1/router"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("hello thanh page")
}

func setupRouters(app *fiber.App) {
	//welcome
	app.Get("/api", welcome)
	//book
	app.Post("/api/books", router.CreateBook)
	app.Get("/api/books", router.Getbooks)
	app.Get("/api/books/:id", router.GetBook)
	app.Delete("/api/books:id", router.DeleteBook)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRouters(app)
	log.Fatal(app.Listen(":4000"))
}
