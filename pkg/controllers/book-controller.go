package controllers

import (
	"book_management_crud_api_sql/pkg/models"
	"book_management_crud_api_sql/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBookModel(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookModel endpoint")

	res, err := json.Marshal(models.Book{})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBook endpoint")

	newBook := &models.Book{}
	if err := utils.ParseBody(r, newBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := newBook.CreateBook(); err != models.OK {
		if err == models.DUPLICATE {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	res, err := json.Marshal(newBook)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllBooks endpoint")

	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookById endpoint")

	bookId := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(bookId, 10, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book := models.GetBookById(id)
	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateBook endpoint")

	updatedBook := &models.Book{}
	if err := utils.ParseBody(r, updatedBook); err != nil || *updatedBook == *new(models.Book) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookId := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(bookId, 10, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := models.UpdateBook(id, updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if book == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBook endpoint")

	bookId := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(bookId, 10, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := models.DeleteBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := json.Marshal(book)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
