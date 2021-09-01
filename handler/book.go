package handler

import (
	"fmt"
	"mvcApi/database"
	"mvcApi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handlerBook struct {
	repo database.RepoBook
}

func NewHandlerBook(repo database.RepoBook) *handlerBook {
	return &handlerBook{repo: repo}
}

func (h *handlerBook) GetBooks(c echo.Context) error {
	books, err := h.repo.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"users":   books,
	})

}

func (h *handlerBook) CreateBook(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	book, err := h.repo.CreateBooks(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func (h *handlerBook) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	book := models.Book{}
	c.Bind(&book)
	newBook, err := h.repo.UpdateBook(id, &book)

	if err != nil {
		fmt.Println("disini", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "oke",
		"updated book": newBook,
	})

}

func (h *handlerBook) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	book := models.Book{}

	if err = h.repo.DeleteBook(id, book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete book in handler",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success delete book with id %d", id),
	})

}

func (h *handlerBook) GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	book, err := h.repo.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get book in handler",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get one book with id %d", id),
		"users":    book,
	})

}
