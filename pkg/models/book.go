package models

import (
	"book_management_crud_api_sql/pkg/config"
	"log"

	"github.com/jinzhu/gorm"
	// "gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type CreateError uint8

const (
	OK CreateError = iota
	DUPLICATE
	DB_ERR
)

func Init() *gorm.DB {
	config.Connect()
	db = config.GetDB()
	if db.AutoMigrate(&Book{}); db.Error != nil {
		log.Println("Failed DB AutoMigrate")
	}
	return db
}

func (b *Book) CreateBook() CreateError {
	log.Printf("CreateBook\n%v\n", b)

	if !db.NewRecord(b) {
		log.Println("Already exists in database")
		return DUPLICATE
	}
	if err := db.Create(b).Error; err != nil {
		log.Println(err)
		return DB_ERR
	}
	log.Println("Added to database")
	return OK
}

func GetAllBooks() []Book {
	log.Println("GetAllBooks")

	var books []Book
	if err := db.Find(&books).Error; err != nil {
		log.Println(err)
	}
	log.Println(books)
	return books
}

func GetBookById(id uint64) (*Book, error) {
	log.Println("GetBookById", id)

	var book Book
	if err := db.Where("ID=?", id).Find(&book).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	if book != *new(Book) {
		log.Println("Found", book)
		return &book, nil
	}
	return nil, nil
}

func UpdateBook(id uint64, newBook *Book) (*Book, error) {
	log.Println("UpdateBook", id)

	var original Book
	book, err := GetBookById(id)
	if book == nil || err != nil {
		return nil, err
	}
	original = *book

	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}

	if *book != original {
		db.Save(&book)
		if db.Error != nil {
			log.Println(err)
			return nil, db.Error
		}
	}
	return book, nil
}

func DeleteBook(id uint64) (*Book, error) {
	log.Println("DeleteBook", id)

	var book *Book
	var err error
	if book, err = GetBookById(id); book == nil || err != nil {
		return nil, err
	}

	if err := db.Where("ID=?", id).Delete(*book).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Deleted", book)
	return book, nil
}
