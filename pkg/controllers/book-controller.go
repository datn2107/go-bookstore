package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/datn2107/go-bookstore/pkg/models"
	"github.com/datn2107/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func parseIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	return ID, err
}

func responseObjectsInfor(w http.ResponseWriter, objects interface{}) {
	res, _ := json.Marshal(objects)

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	responseObjectsInfor(w, newBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	ID, err := parseIdFromRequest(r)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	responseObjectsInfor(w, bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	bookDetails := newBook.CreateBook()
	responseObjectsInfor(w, bookDetails)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	ID, err := parseIdFromRequest(r)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	bookDetails := models.DeleteBook(ID)
	responseObjectsInfor(w, bookDetails)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	ID, err := parseIdFromRequest(r)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	responseObjectsInfor(w, bookDetails)
}
