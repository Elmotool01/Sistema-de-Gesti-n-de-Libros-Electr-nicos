package models // Se declara el paquete models, que agrupa las estructuras de datos del sistema.

// Libro representa la estructura de un libro electrónico dentro del sistema.
// En Go usamos "struct" (estructura) en lugar de clases como en Java.
type Libro struct {
	// ID almacena el identificador único del libro (clave primaria en MySQL).
	ID int

	// Titulo almacena el nombre o título del libro electrónico.
	Titulo string

	// Autor almacena el nombre del autor del libro.
	Autor string

	// Categoria almacena la categoría o género del libro (ej. Programación, Novela).
	Categoria string

	// AnioPublicacion almacena el año de publicación del libro.
	AnioPublicacion int

	// Formato almacena el formato del archivo (PDF, EPUB o MOBI).
	Formato string

	// StockLicencias almacena la cantidad disponible/licencias del libro.
	StockLicencias int
}
