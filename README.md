Este proyecto es un crud de libros, en el podras agregar, actualizar, visualizar, eliminar y buscar por id mediante rutas
Para iniciar este proyecto promero debemos tener golang istalado en nuestras computadoras, puedes ingresar desde esta url https://go.dev/
depsues, tendremos que realizar lo sigueinte, copia el url de este proyecto y en una carpeta de tu gusto realiza el siguiente comando
git clone https://github.com/AMOZINGAS/golang-practice-ma.git
una vez clonado debemos de ir la carpeta donde se encuentra golang-practice-ma
ya estando dentro de la carpeta golang-practice-ma, instalaremos lo siguiente
go install "modernc.org/sqlite"
go mod tidy
go build
go run main.go
con esto podremos ver el mensaje de que nuestro servidor esta escuchando en http://localhost:8080
para poder testear el crud es recomendado ir a postamn y ejecutar lo siguiente para cada ruta
Create 
http://localhost:8080/books
iremos a row y seleccionaremos json
y pondremos lo siguiente
{ 
"titulo": "el titulo de tu libro favorito",
"autor": "autor de tu libro favorito"
}
y le damos en enviar
si quieres agregar mas libros solo debes cambiar los nombres o el titulo y seleccionar enviar de nuevo
una vez que los libros hayan sido enviados, los podemos visualizar en postman asi
Get (todos)
http://localhost:8080/books
Get (por id)
http://localhost:8080/books/id (es el id del libro que deseas encontrar)

para actualizar seleccionamos 
Put
http://localhost:8080/books/id (es el id del libro que deseas modificar)
seleccionamos de nuevo row y json
{ 
"titulo": "titulo nuevo",
"autor": "autor nuevo"
}

y finalmente para el delete es de la siguiente manera
Delete
http://localhost:8080/books/id (es el id del libro que deseas eliminar)
y seleccionamos enviar





