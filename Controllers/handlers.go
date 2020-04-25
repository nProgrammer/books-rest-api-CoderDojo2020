package Controllers

import (
	"books-rest-api/Repositories"
	"books-rest-api/Utils"
	"database/sql"
    "encoding/json"
    "net/http"
	"books-rest-api/Models"
	"strconv"
	"github.com/gorilla/mux"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var error Models.Error
    	var book Models.Book
        var books []Models.Book

        books = []Models.Book{}
        books, err := Repositories.GetBooksRep(db, book, books)

        if err != nil {
        	error.Message = "Server error"
        	Utils.SendError(w, http.StatusInternalServerError, error)
		}
		w.Header().Set("Content-Type","application/json")
		Utils.SendSucces(w, books)
    }
}

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book Models.Book
		var bookID int
		var error Models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields."
			Utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookID, err := Repositories.CreateBookRep(db, book)

		if err != nil {
			error.Message = "Server error"
			Utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type","text/plain")
		Utils.SendSucces(w, bookID)

		json.NewEncoder(w).Encode(book.ID)
	}
}

func UpdateBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var error Models.Error
    	var book Models.Book

        json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields."
			Utils.SendError(w, http.StatusBadRequest, error)
			return
		}

        rowsUpdated, err := Repositories.UpdateBookRep(db, book)

        if err != nil {
        	error.Message = "Server error"
        	Utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

        w.Header().Set("Content-Type","text/plain")
        Utils.SendSucces(w, rowsUpdated)
    }
}

func DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error Models.Error
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := Repositories.RemoveBookRep(db, id)
		if err != nil {
			error.Message = "Server error."
			Utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type","text/plain")
		Utils.SendSucces(w, rowsDeleted)
	}
}