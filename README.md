# Sistema-de-Gesti-n-de-Libros-Electr-nicos-e-Books-
PROGRAMACIÓN ORIENTADA A OBJETOS
Autor: Herrera Cabezas Camilo Nicolas
Fecha: 17/11/2025
# [E-Books]

## Descripción del Proyecto
Este proyecto consiste en el desarrollo de un **sistema básico de gestión de libros electrónicos**, implementado en el lenguaje de programación **Go (Golang)**, que permite registrar, listar y buscar libros a través de un **menú interactivo por consola**.

El sistema utiliza una **base de datos SQL (SQLite)** para garantizar la **persistencia de la información**, asegurando que los datos no se pierdan al cerrar el programa.

---

## Objetivo del Proyecto
Desarrollar un sistema funcional que aplique conceptos de **Programación Orientada a Objetos**, como encapsulación e interfaces, junto con el uso de una **base de datos SQL**, permitiendo la gestión básica de libros electrónicos de forma persistente.

---

## Funcionalidades Principales (8)

1. **Registro de libros electrónicos**  
   Permite ingresar nuevos libros proporcionando título, autor y año de publicación.

2. **Validación de datos de entrada**  
   El sistema verifica que los campos obligatorios no estén vacíos y que el año sea válido antes de guardar la información.

3. **Asignación automática de identificador (ID)**  
   Cada libro registrado obtiene un ID único generado automáticamente por la base de datos.

4. **Listado de libros registrados**  
   Muestra todos los libros almacenados en la base de datos, recuperando la información mediante consultas SQL.

5. **Búsqueda de libros por ID**  
   Permite localizar un libro específico utilizando su identificador único.

6. **Persistencia de datos con SQL**  
   La información se almacena en una base de datos SQLite (`ebooks.db`), manteniendo los datos incluso después de cerrar el programa.

7. **Encapsulación del acceso a la base de datos**  
   El acceso a SQL se maneja a través de una implementación concreta (`SQLLibrary`), evitando el acceso directo desde el flujo principal del programa.

8. **Interfaz de usuario por consola**  
   El sistema ofrece un menú claro e intuitivo que guía al usuario en la interacción con el programa.

---

## Tecnologías Utilizadas

- **Lenguaje de programación:** Go (Golang)
- **Base de datos:** SQLite (SQL)
- **Librerías estándar de Go:**
  - `bufio` – Lectura de datos por consola
  - `fmt` – Entrada y salida de información
  - `os` – Manejo del sistema
  - `strconv` – Conversión de tipos
  - `strings` – Manipulación de cadenas
  - `database/sql` – Gestión de bases de datos SQL
- **Driver SQLite:** `modernc.org/sqlite`
- **Entorno de desarrollo:** Visual Studio Code
- **Sistema de control de versiones:** Git / GitHub

## Tecnologías usadas
- Backend: [Ej: Node.js + Express / Python Flask / Java Spring]
- Base de datos: [Ej: SQLite / PostgreSQL / MongoDB]
- Documentación: [Ej: Swagger / Postman]
- Control de versiones: Git + GitHub

