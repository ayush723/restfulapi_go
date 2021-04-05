package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholde for book
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Panicln(err)
	}
	//create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books(id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfuly created table books!")
	}
	statement.Exec()
	// create
	statement, _ = db.Prepare("INSERT INTO books(name, author, isbn) VALUES(?,?,?)")
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("inserted the book into database!")
	// read
	rows, _ := db.Query("Select id, name, author FORM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// update
	statement, _ = db.Prepare("update books set name = ? where id =?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated yhe book in database!")
	// delete
	statement, _ = db.Prepare("delete from books where id =?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")

}