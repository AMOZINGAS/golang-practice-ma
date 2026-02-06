package main

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"fmt"
	"golang-practice-ma/internal/service"
	"golang-practice-ma/internal/store"
	"golang-practice-ma/internal/transport"
	"log"
	"net/http"
)

func main() {

	//conectar a SQLLite
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Crear el table si no existe
	q := `CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	)`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	//Inyectar nuestras dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	//configurar las rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBooksByID)

	fmt.Println("Servidor ejecutandose en http://localhost:8080")

	fmt.Println("api endpoints")
	fmt.Println("get /books 			-Obtener todos los libros")
	fmt.Println("post /books 			-Crear un nuevo libro")
	fmt.Println("get /books/{id} 		-Obtener un libro especifico")
	fmt.Println("put /books/{id} 		-Actualizar un libro")
	fmt.Println("delete /books/{id} 	-Eliminar un libro")

	//empezar y escuchar al servidor
	log.Fatal(http.ListenAndServe(":8080", nil))

}
