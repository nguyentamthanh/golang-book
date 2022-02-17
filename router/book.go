package router

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/thanh/go-book1/database"
	"github.com/thanh/go-book1/model"
)

type Book struct {
	Id     uint   `json:"id" `
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
	Year   int    `json:"year"`
}

func CreateResponseBook(book model.Book) Book {
	return Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
		Rating: book.Rating,
		Year:   book.Year,
	}
}

func CreateBook(c *fiber.Ctx) error {
	var book model.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&book)
	responseBook := CreateResponseBook(book)
	return c.Status(200).JSON(responseBook)
}

func Getbooks(c *fiber.Ctx) error {
	books := []model.Book{}
	database.Database.Db.Find(&books)
	responseBooks := []Book{}
	for _, book := range books {
		responseBook := CreateResponseBook(book)
		responseBooks = append(responseBooks, responseBook)
	}
	return c.Status(200).JSON(responseBooks)
}

func findBook(id int, book *model.Book) error {
	database.Database.Db.Find(&book, "id = ?", id)
	if book.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var book model.Book

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findBook(id, &book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseBook(book)

	return c.Status(200).JSON(responseUser)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var book model.Book
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	err = findBook(id, &book)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Rating int    `json:"rating"`
		Year   int    `json:"year"`
	}
	var updateData UpdateBook
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	book.Title = updateData.Title
	book.Rating = updateData.Rating
	book.Author = updateData.Author
	book.Year = updateData.Year

	database.Database.Db.Save((&book))
	responseBook := CreateResponseBook(book)
	return c.Status(200).JSON(responseBook)
}
func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var book model.Book

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findBook(id, &book)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&book).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}
