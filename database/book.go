package database

import (
	"errors"
	"mvcApi/models"

	"gorm.io/gorm"
)

type repoBook struct {
	db *gorm.DB
}

type RepoBook interface {
	GetBooks() ([]models.Book, error)
	CreateBooks(book models.Book) (models.Book, error)
	DeleteBook(id int, book models.Book) error
	UpdateBook(id int, NewBook *models.Book) (interface{}, error)
	GetBookByID(id int) (models.Book, error)
}

func NewRepoBook(db *gorm.DB) *repoBook {
	return &repoBook{db: db}
}

func (r *repoBook) GetBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repoBook) CreateBooks(book models.Book) (models.Book, error) {

	if err := r.db.Save(&book).Error; err != nil {
		return models.Book{}, errors.New("failed to create book")
	}

	return book, nil

}

func (r *repoBook) DeleteBook(id int, book models.Book) error {

	err := r.db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (r *repoBook) UpdateBook(id int, NewBook *models.Book) (interface{}, error) {

	book := models.Book{}
	tx := r.db.Find(&book, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected > 0 {
		tx = r.db.Model(&book).Updates(NewBook)

		if tx.Error != nil {
			return nil, tx.Error
		}

		return book, nil
	}

	return nil, tx.Error
}

func (r *repoBook) GetBookByID(id int) (models.Book, error) {
	var book models.Book

	if err := r.db.Find(&book, id).Error; err != nil {
		return models.Book{}, errors.New("cant find user with id")
	}

	return book, nil
}
