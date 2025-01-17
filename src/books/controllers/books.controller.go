package booksController

import (
	"fiber-app/src/books/dtos"
	"fiber-app/src/books/services"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	bookService *bookService.BookService
}

func NewBookController() *BookController {
	return &BookController{
		bookService: bookService.NewBookService(),
	}
}

func (bc *BookController) GetBooks(c *fiber.Ctx) error {
	books, err := bc.bookService.GetAllBooks(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data": books})
}

func (bc *BookController) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	book, err := bc.bookService.GetBookByID(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": book})
}

func (bc *BookController) CreateBook(c *fiber.Ctx) error {
	b := new(dtos.CreateDTO) // Use createDTO
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	result, err := bc.bookService.CreateBook(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create book", "message": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"result": result})
}

func (bc *BookController) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	b := new(dtos.UpdateDTO) // Use updateDTO
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	result, err := bc.bookService.UpdateBook(c.Context(), id, b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update book", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"result": result})
}

func (bc *BookController) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	result, err := bc.bookService.DeleteBook(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete book", "message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"result": result})
}
