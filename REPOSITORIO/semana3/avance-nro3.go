/*
@titulo: programa principal del proyecto
@autor: Camilo Nicolas Herrera Cabezas
@fecha: 21/11/2025
@descripcion: Sistema básico de gestión de libros electrónicos con encapsulación, interfaces y manejo de errores.
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

type Library interface {
	AddBook(title, author string, year int) (Book, error)
	ListBooks() []Book
	FindBookByID(id int) (*Book, error)
}

type SQLLibrary struct {
	db *sql.DB
}

func NewSQLLibrary(db *sql.DB) *SQLLibrary {
	return &SQLLibrary{db: db}
}

func OpenDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:"+dbPath)
	if err != nil {
		return nil, err
	}
	schema := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		year INTEGER NOT NULL
	);`
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (l *SQLLibrary) AddBook(title, author string, year int) (Book, error) {
	res, err := l.db.Exec(
		"INSERT INTO books(title, author, year) VALUES(?,?,?)",
		title, author, year,
	)
	if err != nil {
		return Book{}, err
	}
	id, _ := res.LastInsertId()
	return Book{ID: int(id), Title: title, Author: author, Year: year}, nil
}

func (l *SQLLibrary) ListBooks() []Book {
	rows, err := l.db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return []Book{}
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year)
		books = append(books, b)
	}
	return books
}

func (l *SQLLibrary) FindBookByID(id int) (*Book, error) {
	var b Book
	err := l.db.QueryRow(
		"SELECT id, title, author, year FROM books WHERE id=?",
		id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.Year)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	db, err := OpenDB("ebooks.db")
	if err != nil {
		fmt.Println("Error BD:", err)
		return
	}
	defer db.Close()

	var library Library = NewSQLLibrary(db)

	for {
		fmt.Println("\n1. Registrar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar por ID")
		fmt.Println("4. Salir")

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		switch op {
		case "1":
			fmt.Print("Título: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Autor: ")
			author, _ := reader.ReadString('\n')
			fmt.Print("Año: ")
			y, _ := reader.ReadString('\n')
			year, _ := strconv.Atoi(strings.TrimSpace(y))

			book, _ := library.AddBook(
				strings.TrimSpace(title),
				strings.TrimSpace(author),
				year,
			)
			fmt.Println("Libro guardado con ID:", book.ID)

		case "2":
			for _, b := range library.ListBooks() {
				fmt.Printf("%d | %s | %s | %d\n", b.ID, b.Title, b.Author, b.Year)
			}

		case "3":
			fmt.Print("ID: ")
			i, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(i))
			b, err := library.FindBookByID(id)
			if err != nil {
				fmt.Println("No encontrado")
			} else {
				fmt.Printf("%d | %s | %s | %d\n", b.ID, b.Title, b.Author, b.Year)
			}

		case "4":
			return
		}
	}
}
