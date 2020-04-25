package main

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"net/http"
	"encoding/json"
	"database/sql"
	"fmt"
	"books-rest-api/Drivers"
	"books-rest-api/Controllers"
)

var db *sql.DB 

func init() {
	gotenv.Load()
}

func main() {
	db = Drivers.ConnectDB()
	fmt.Println(db)

	r := mux.NewRouter()

	r.HandleFunc("/", hello).Methods("GET")  // CRUD
	r.HandleFunc("/getBooks", Controllers.GetBooks(db)).Methods("GET") 
	r.HandleFunc("/addBook", Controllers.CreateBook(db)).Methods("POST")
	r.HandleFunc("/updateBook", Controllers.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/deleteBook/", Controllers.DeleteBook(db)).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}