/*
@titulo: programa principal del proyecto
@autor: Camilo Nicolas Herrera Cabezas
@fecha: 16/11/2025
@descripcion: Sistema básico de gestión de libros electrónicos con encapsulación, interfaces y manejo de errores.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/////////////////////////////////////////////////////
// "CLASES" (STRUCTS) E INTERFACES DEL SISTEMA
/////////////////////////////////////////////////////

// Book representa un libro electrónico registrado en el sistema.
// Esta "clase" modela los datos esenciales de un libro.
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// Library define el contrato (INTERFAZ) para cualquier tipo de biblioteca.
// De esta forma, el código principal trabaja contra la interfaz y no
// contra una implementación específica (principio de abstracción).
type Library interface {
	AddBook(title, author string, year int) (Book, error)
	ListBooks() []Book
	FindBookByID(id int) (*Book, error)
}

// InMemoryLibrary es una implementación concreta de Library.
// Aplica ENCAPSULACIÓN porque guarda internamente la colección de libros
// y el siguiente ID a asignar; el resto del programa solo accede a través
// de los métodos definidos por la interfaz.
type InMemoryLibrary struct {
	books  []Book
	nextID int
}

// NewInMemoryLibrary es un constructor que crea una biblioteca en memoria.
func NewInMemoryLibrary() *InMemoryLibrary {
	return &InMemoryLibrary{
		books:  []Book{},
		nextID: 1,
	}
}

/////////////////////////////////////////////////////
// MÉTODOS DE InMemoryLibrary (ENCAPSULACIÓN + ERRORES)
/////////////////////////////////////////////////////

// AddBook agrega un nuevo libro a la biblioteca.
//
// Esta funcionalidad:
// 1) Valida los datos de entrada (manejo de errores).
// 2) Encapsula la generación del ID (nextID).
// 3) Devuelve el libro creado o un error si algo va mal.
func (l *InMemoryLibrary) AddBook(title, author string, year int) (Book, error) {
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	if title == "" || author == "" {
		return Book{}, fmt.Errorf("el título y el autor son obligatorios")
	}
	if year <= 0 {
		return Book{}, fmt.Errorf("el año debe ser mayor que cero")
	}

	newBook := Book{
		ID:     l.nextID,
		Title:  title,
		Author: author,
		Year:   year,
	}

	l.books = append(l.books, newBook)
	l.nextID++

	return newBook, nil
}

// ListBooks devuelve la lista de libros almacenados.
func (l *InMemoryLibrary) ListBooks() []Book {
	return l.books
}

// FindBookByID busca un libro por su ID.
//
// Esta funcionalidad recorre la colección interna y devuelve:
// - Un puntero al libro encontrado.
// - O un error en caso de que no exista.
func (l *InMemoryLibrary) FindBookByID(id int) (*Book, error) {
	for i := range l.books {
		if l.books[i].ID == id {
			return &l.books[i], nil
		}
	}
	return nil, fmt.Errorf("no existe un libro con el ID %d", id)
}

/////////////////////////////////////////////////////
// FUNCIONES DE APOYO PARA INTERACCIÓN CON EL USUARIO
/////////////////////////////////////////////////////

// ReadLine lee una línea de texto desde consola mostrando un mensaje previo.
// Devuelve la cadena leída o un error si algo falla en la lectura.
func ReadLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		// Manejo de error de entrada/salida
		return "", fmt.Errorf("error al leer desde la consola: %w", err)
	}
	text = strings.TrimSpace(text)
	return text, nil
}

// ReadInt lee un número entero desde consola con validación.
// Esta función implementa manejo de errores de conversión (string -> int)
// y obliga al usuario a ingresar un número válido.
func ReadInt(reader *bufio.Reader, prompt string) (int, error) {
	for {
		input, err := ReadLine(reader, prompt)
		if err != nil {
			// Error de lectura desde la consola
			return 0, err
		}

		value, convErr := strconv.Atoi(input)
		if convErr != nil {
			fmt.Println("⚠ Entrada no válida. Por favor, ingrese un número entero.")
			continue
		}
		return value, nil
	}
}

// PrintBooks imprime en consola un listado de libros.
// Se separa esta lógica para mantener el main más limpio.
func PrintBooks(books []Book) {
	if len(books) == 0 {
		fmt.Println("No hay libros registrados todavía.")
		return
	}

	fmt.Println("\n=== Listado de libros electrónicos registrados ===")
	for _, book := range books {
		fmt.Printf("ID: %d | Título: %s | Autor: %s | Año: %d\n",
			book.ID, book.Title, book.Author, book.Year)
	}
}

/////////////////////////////////////////////////////
// FUNCIÓN PRINCIPAL (MENÚ + FLUJO DE CONTROL)
/////////////////////////////////////////////////////

func main() {
	// Punto de entrada del programa.
	reader := bufio.NewReader(os.Stdin)

	// Trabajamos contra la INTERFAZ Library, pero usamos
	// la implementación concreta InMemoryLibrary.
	var library Library = NewInMemoryLibrary()

	fmt.Println("=== Sistema de Gestión de Libros Electrónicos ===")

	for {
		fmt.Println("\nMenú principal:")
		fmt.Println("1. Registrar un nuevo libro")
		fmt.Println("2. Listar libros registrados")
		fmt.Println("3. Buscar libro por ID")
		fmt.Println("4. Salir")

		option, err := ReadLine(reader, "Seleccione una opción (1-4): ")
		if err != nil {
			fmt.Println("Error al leer la opción del menú:", err)
			continue
		}

		switch option {
		case "1":
			// -------------------------------
			// REGISTRO DE UN NUEVO LIBRO
			// -------------------------------
			fmt.Println("\n=== Registro de un nuevo libro electrónico ===")

			title, err := ReadLine(reader, "Ingrese el título del libro: ")
			if err != nil {
				fmt.Println("Error al leer el título:", err)
				continue
			}

			author, err := ReadLine(reader, "Ingrese el autor del libro: ")
			if err != nil {
				fmt.Println("Error al leer el autor:", err)
				continue
			}

			year, err := ReadInt(reader, "Ingrese el año de publicación: ")
			if err != nil {
				fmt.Println("Error al leer el año de publicación:", err)
				continue
			}

			// Aquí usamos la funcionalidad encapsulada en AddBook,
			// que valida los datos y genera el ID automáticamente.
			newBook, err := library.AddBook(title, author, year)
			if err != nil {
				fmt.Println("No se pudo registrar el libro:", err)
				continue
			}

			fmt.Printf("✔ Libro registrado correctamente con ID %d.\n", newBook.ID)

		case "2":
			// -------------------------------
			// LISTADO DE LIBROS
			// -------------------------------
			books := library.ListBooks()
			PrintBooks(books)

		case "3":
			// -------------------------------
			// BÚSQUEDA DE LIBRO POR ID
			// -------------------------------
			id, err := ReadInt(reader, "Ingrese el ID del libro a buscar: ")
			if err != nil {
				fmt.Println("Error al leer el ID:", err)
				continue
			}

			book, err := library.FindBookByID(id)
			if err != nil {
				fmt.Println("Resultado de la búsqueda:", err)
				continue
			}

			fmt.Println("\nLibro encontrado:")
			fmt.Printf("ID: %d | Título: %s | Autor: %s | Año: %d\n",
				book.ID, book.Title, book.Author, book.Year)

		case "4":
			fmt.Println("\nFin del programa. ¡Gracias por usar el sistema!")
			return

		default:
			fmt.Println("⚠ Opción no válida. Intente nuevamente.")
		}
	}
}
