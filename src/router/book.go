package router

import (
	booksController "fiber-app/src/books/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddBookGroup(app *fiber.App) {
		// Initialize the controller and routes
	bookController := booksController.NewBookController()
	bookGroup := app.Group("/books")

	// Add route handlers from the booksController
	bookGroup.Get("/", bookController.GetBooks)         // Fetch all books
	bookGroup.Get("/:id", bookController.GetBook)       // Fetch a specific book by ID
	bookGroup.Post("/", bookController.CreateBook)      // Create a new book
	bookGroup.Put("/:id", bookController.UpdateBook)    // Update a book by ID
	bookGroup.Delete("/:id", bookController.DeleteBook) // Delete a book by ID
}
