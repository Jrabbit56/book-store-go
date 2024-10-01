package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type BookHandler struct {
	service ports.BookService
}

func NewBookHandler(service ports.BookService) *BookHandler {
	return &BookHandler{service: service}
}

type SwaggerFiberMap map[string]interface{}

// GetAllBooks godoc
//
//	@Summary		Get all books
//	@Description	Retrieve a list of all books
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.SwaggerBook
//	@Failure		500	{object}	handlers.SwaggerFiberMap
//	@Router			/books [get]
func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

// GetBook godoc
//
//	@Summary		Get a book by ID
//	@Description	Retrieve a book's details using its ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	domain.SwaggerBook
//	@Failure		400	{object}	map[string]string	"error"
//	@Failure		404	{object}	map[string]string	"error"
//	@Router			/books/{id} [get]
func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	book, err := h.service.GetBook(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

// CreateBook godoc
//
//	@Summary		Create a new book
//	@Description	Create a new book with the provided details
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.SwaggerBook			true	"Book object"
//	@Success		200		{object}	map[string]string	"message"
//	@Failure		400		{object}	map[string]string	"error"
//	@Failure		500		{object}	map[string]string	"error"
//	@Router			/books [post]
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := new(domain.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.CreateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	//return c.Status(fiber.StatusCreated).JSON(book)
	return c.JSON(fiber.Map{"message": "Create Book Successfully!!"})
}

// UpdateBook godoc
//
//	@Summary		Update a book
//	@Description	Updates an existing book with the given ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Book ID"
//	@Param			book	body		domain.SwaggerBook	true	"Updated book information"
//	@Success		200		{object}	handlers.SwaggerFiberMap
//	@Failure		400		{object}	handlers.SwaggerFiberMap
//	@Failure		500		{object}	handlers.SwaggerFiberMap
//	@Router			/books/{id} [put]
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	book := new(domain.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body: " + err.Error()})
	}
	book.ID = uint(id)

	if err := h.service.UpdateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// return c.JSON(book)
	return c.JSON(fiber.Map{"message": "Update Successfully!!"})
}

// DeleteBook godoc
//
//	@Summary		Delete a book
//	@Description	Deletes a book with the given ID
//	@Tags			books
//	@Produce		json
//	@Param			id	path	int	true	"Book ID"
//	@Success		204	"No Content"
//	@Failure		400	{object}	handlers.SwaggerFiberMap
//	@Failure		500	{object}	handlers.SwaggerFiberMap
//	@Router			/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
