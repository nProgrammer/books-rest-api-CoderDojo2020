package Repositories

import (
	"books-rest-api/Models"
	"database/sql"
)

func errIF(err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return 0, err
}

func GetBooksRep(db *sql.DB, book Models.Book, books []Models.Book) ([]Models.Book, error) {
	rows, err := db.Query("select * from books")

	if err != nil {
		return []Models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []Models.Book{}, err
	}

	return books, err
}

func CreateBookRep(db *sql.DB, book Models.Book) (int, error) {
	err := db.QueryRow("insert into books (Title, Author, Year) values ($1, $2, $3) returning ID;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	errIF(err)

	return book.ID, nil
}

func UpdateBookRep(db *sql.DB, book Models.Book) (int64, error) {
	result, err := db.Exec("update books set Title=$1, Author=$2, Year=$3 where ID=$4 returning ID",
		&book.Title, &book.Author, &book.Year, &book.ID)

	errIF(err)

	rowsUpdated, err := result.RowsAffected()

	errIF(err)

	return rowsUpdated, nil
}

func RemoveBookRep(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where ID=$1", id)
	errIF(err)
	rowsDeleted, err := result.RowsAffected()
	errIF(err)
	return rowsDeleted, nil
}