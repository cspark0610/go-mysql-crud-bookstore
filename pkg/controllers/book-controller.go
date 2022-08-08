package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cspark0610/go-mysql-bookstore/pkg/models"
	"github.com/cspark0610/go-mysql-bookstore/pkg/utils"

	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["bookId"], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookById(ID)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// equivalente al res.send(res) en nodejs
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["bookId"], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	//1.getBookById 2. override each field, and 3. save
	book, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
