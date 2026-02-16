package main // Paquete principal.

import ( // Imports necesarios.
	"bufio" // Lectura por consola.
	"fmt"   // Impresión por consola.
	"os"    // Acceso a stdin.
)

// main es el punto de entrada del programa.
func main() {
	// Creamos el lector para capturar entradas del usuario.
	reader := bufio.NewReader(os.Stdin)

	// Creamos un repositorio persistente en JSON (guarda/carga desde libros.json).
	repoJSON, err := NewRepoJSON("libros.json")
	if err != nil {
		fmt.Println("❌ Error al inicializar el repositorio JSON:", err)
		return
	}

	// Usamos la interfaz para desacoplar main de la implementación.
	var repo RepositorioLibros = repoJSON

	// Cargamos automáticamente 20 libros reales SOLO si aún no hay libros guardados.
	// Esto evita que se dupliquen cada vez que inicias el programa.
	precargarSiVacio(repo)

	// Bucle principal del menú.
	for {
		fmt.Println("\n=== SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS ===")
		fmt.Println("1) Agregar libro")
		fmt.Println("2) Listar libros")
		fmt.Println("3) Buscar libro por ID")
		fmt.Println("4) Actualizar libro")
		fmt.Println("5) Eliminar libro")
		fmt.Println("0) Salir")
		fmt.Print("Opción: ")

		opcion := leerLinea(reader)

		switch opcion {
		case "1":
			agregarLibro(reader, repo)
		case "2":
			listarLibros(repo)
		case "3":
			buscarLibro(reader, repo)
		case "4":
			actualizarLibro(reader, repo)
		case "5":
			eliminarLibro(reader, repo)
		case "0":
			fmt.Println("Saliendo... ✅")
			return
		default:
			fmt.Println("❌ Opción inválida.")
		}
	}
}

// precargarSiVacio verifica si el repositorio está vacío; si lo está, inserta 20 libros “reales”.
func precargarSiVacio(repo RepositorioLibros) {
	// Si ya hay libros (porque vienen de libros.json), no hacemos nada.
	if len(repo.Listar()) > 0 {
		return
	}

	// Lista de 20 libros realistas.
	datos := []struct {
		id     int
		titulo string
		autor  string
		anio   int
	}{
		{1, "Clean Code", "Robert C. Martin", 2008},
		{2, "The Pragmatic Programmer", "Andrew Hunt & David Thomas", 1999},
		{3, "Design Patterns", "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides", 1994},
		{4, "Refactoring", "Martin Fowler", 1999},
		{5, "Working Effectively with Legacy Code", "Michael C. Feathers", 2004},
		{6, "Introduction to Algorithms", "Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, Clifford Stein", 2009},
		{7, "Code Complete", "Steve McConnell", 2004},
		{8, "The Clean Coder", "Robert C. Martin", 2011},
		{9, "Effective Java", "Joshua Bloch", 2018},
		{10, "Head First Design Patterns", "Eric Freeman & Elisabeth Robson", 2004},
		{11, "The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", 2015},
		{12, "Pro Git", "Scott Chacon & Ben Straub", 2014},
		{13, "Artificial Intelligence: A Modern Approach", "Stuart Russell & Peter Norvig", 2021},
		{14, "Structure and Interpretation of Computer Programs", "Harold Abelson & Gerald Jay Sussman", 1996},
		{15, "Computer Networks", "Andrew S. Tanenbaum & David J. Wetherall", 2010},
		{16, "Operating System Concepts", "Abraham Silberschatz, Peter B. Galvin, Greg Gagne", 2018},
		{17, "Compilers: Principles, Techniques, and Tools", "Alfred V. Aho, Monica S. Lam, Ravi Sethi, Jeffrey D. Ullman", 2006},
		{18, "Security Engineering", "Ross Anderson", 2020},
		{19, "The Art of Computer Programming", "Donald E. Knuth", 2011},
		{20, "The Mythical Man-Month", "Frederick P. Brooks Jr.", 1995},
	}

	// Insertamos cada libro usando el constructor (validación) y el repo (que guarda en JSON).
	for _, d := range datos {
		libPtr, err := NewLibro(d.id, d.titulo, d.autor, d.anio)
		if err != nil {
			continue
		}
		_ = repo.Agregar(*libPtr)
	}

	fmt.Println("✅ Se precargaron 20 libros y quedaron guardados en libros.json.")
}

/* =======================
   FUNCIONES CRUD (MENÚ)
   ======================= */

// agregarLibro solicita datos, crea un Libro validado y lo guarda.
func agregarLibro(reader *bufio.Reader, repo RepositorioLibros) {
	id := leerEntero(reader, "ID: ")
	fmt.Print("Título: ")
	titulo := leerLinea(reader)
	fmt.Print("Autor: ")
	autor := leerLinea(reader)
	anio := leerEntero(reader, "Año: ")

	libPtr, err := NewLibro(id, titulo, autor, anio)
	if err != nil {
		fmt.Println("❌ Error de validación:", err)
		return
	}

	if err := repo.Agregar(*libPtr); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ Libro agregado correctamente.")
}

// listarLibros imprime todos los libros.
func listarLibros(repo RepositorioLibros) {
	libros := repo.Listar()
	if len(libros) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}

	fmt.Println("\n--- LISTA DE LIBROS ---")
	for _, l := range libros {
		fmt.Printf("[%d] %s | %s | %d\n", l.ID(), l.Titulo(), l.Autor(), l.Anio())
	}
}

// buscarLibro busca por ID y muestra el resultado.
func buscarLibro(reader *bufio.Reader, repo RepositorioLibros) {
	id := leerEntero(reader, "ID a buscar: ")

	libro, err := repo.BuscarPorID(id)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Printf("✅ Encontrado: [%d] %s | %s | %d\n", libro.ID(), libro.Titulo(), libro.Autor(), libro.Anio())
}

// actualizarLibro actualiza campos usando setters (encapsulación).
func actualizarLibro(reader *bufio.Reader, repo RepositorioLibros) {
	id := leerEntero(reader, "ID a actualizar: ")

	libro, err := repo.BuscarPorID(id)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Printf("Actual: %s | %s | %d\n", libro.Titulo(), libro.Autor(), libro.Anio())

	fmt.Print("Nuevo título (Enter para mantener): ")
	nuevoTitulo := leerLinea(reader)
	fmt.Print("Nuevo autor (Enter para mantener): ")
	nuevoAutor := leerLinea(reader)
	fmt.Print("Nuevo año (Enter para mantener, o número): ")
	nuevoAnioStr := leerLinea(reader)

	libMod := libro

	if nuevoTitulo != "" {
		if err := (&libMod).SetTitulo(nuevoTitulo); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
	}

	if nuevoAutor != "" {
		if err := (&libMod).SetAutor(nuevoAutor); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
	}

	if nuevoAnioStr != "" {
		var nuevoAnio int
		_, err := fmt.Sscanf(nuevoAnioStr, "%d", &nuevoAnio)
		if err != nil {
			fmt.Println("❌ Año inválido.")
			return
		}
		if err := (&libMod).SetAnio(nuevoAnio); err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
	}

	if err := repo.Actualizar(libMod); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ Libro actualizado correctamente.")
}

// eliminarLibro elimina por ID.
func eliminarLibro(reader *bufio.Reader, repo RepositorioLibros) {
	id := leerEntero(reader, "ID a eliminar: ")

	if err := repo.Eliminar(id); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ Libro eliminado correctamente.")
}
