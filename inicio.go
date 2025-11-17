/*
@titulo: programa principal dle proyecto
@autor: Camilo Nicolas Herrera Cabezas
@fecha: 16/11/2025
@descripcion: este programa presenta los aspectos iniciales de go
*/

package main

// Definimos que este archivo pertenece al paquete "main", que es el punto de entrada de un programa en Go.

// Importamos "bufio" para poder leer texto desde la consola de forma cómoda (entrada del usuario).

// Importamos "fmt" para poder imprimir mensajes en la consola y formatear texto.

// Importamos "os" para acceder a funcionalidades del sistema operativo, como la entrada estándar (teclado).

type Book struct {
	// Definimos un nuevo tipo llamado "Book" (Libro) usando una estructura (struct).

	ID int
	// Campo "ID" para identificar de forma única cada libro con un número entero.

	Title string
	// Campo "Title" para almacenar el título del libro como una cadena de texto.

	Author string
	// Campo "Author" para almacenar el nombre del autor del libro.

	Year int
	// Campo "Year" para almacenar el año de publicación del libro.
}
