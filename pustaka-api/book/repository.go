package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindOne(ID int) (Book, error)
	Craete(book Book) (Book, error)
	Update(ID int, book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]Book, error) {
	var books []Book
	err := repo.db.Find(&books).Error
	return books, err
}

func (repo *repository) FindOne(ID int) (Book, error) {
	var book Book
	err := repo.db.Debug().Where("id = ?", ID).First(&book).Error
	return book, err
}

func (repo *repository) Create(book Book) (Book, error) {
	err := repo.db.Create(&book).Error
	return book, err
}

func (repo *repository) Update(ID int, book Book) (Book, error) {
	book, _ = repo.FindOne(ID)
	book.Title = "Niana {Revised Edition}"
	err := repo.db.Save(&book).Error
	return book, err
}
