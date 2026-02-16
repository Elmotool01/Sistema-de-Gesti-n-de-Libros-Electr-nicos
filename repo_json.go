package main // Paquete principal.

import ( // Inicio de imports.
	"encoding/json" // Para convertir a/desde JSON.
	"os"            // Para leer/escribir archivos.
) // Fin de imports.

// LibroDTO es una estructura “exportable” para JSON (campos públicos).
// Se usa para no exponer directamente los campos privados de Libro en el archivo.
type LibroDTO struct { // Define el formato que se guardará en JSON.
	ID     int    `json:"id"`     // Campo público ID para JSON.
	Titulo string `json:"titulo"` // Campo público Titulo para JSON.
	Autor  string `json:"autor"`  // Campo público Autor para JSON.
	Anio   int    `json:"anio"`   // Campo público Anio para JSON.
} // Fin de la estructura DTO.

// RepoJSON implementa RepositorioLibros guardando los datos en un archivo JSON.
type RepoJSON struct { // Definimos el repositorio persistente.
	archivo string       // Ruta/nombre del archivo JSON.
	mem     *RepoMemoria // Repositorio en memoria (map + slice) usado internamente.
} // Fin de RepoJSON.

// NewRepoJSON crea el repositorio JSON, carga datos si el archivo existe y deja listo el repo.
func NewRepoJSON(archivo string) (*RepoJSON, error) { // Constructor del repositorio JSON.
	r := &RepoJSON{ // Crea una instancia del repo JSON.
		archivo: archivo,          // Guarda el nombre/ruta del archivo.
		mem:     NewRepoMemoria(), // Inicializa el repositorio en memoria.
	} // Fin de la instancia.

	_ = r.cargar() // Intenta cargar (si no existe, no pasa nada).
	return r, nil  // Devuelve el repositorio listo.
} // Fin de NewRepoJSON.

// cargar lee el archivo JSON y lo carga al repositorio en memoria.
func (r *RepoJSON) cargar() error { // Función para cargar desde disco.
	data, err := os.ReadFile(r.archivo) // Lee todo el archivo.
	if err != nil {                     // Si ocurrió error al leer...
		if os.IsNotExist(err) { // Si el archivo no existe...
			return nil // No es un problema: empezamos vacío.
		}
		return err // Si es otro error, lo devolvemos.
	}

	var lista []LibroDTO                                 // Crea una lista para deserializar JSON.
	if err := json.Unmarshal(data, &lista); err != nil { // Convierte JSON a estructura.
		return err // Si el JSON está dañado, retorna error.
	}

	for _, d := range lista { // Recorre cada elemento del JSON.
		libPtr, err := NewLibro(d.ID, d.Titulo, d.Autor, d.Anio) // Crea Libro validando.
		if err != nil {                                          // Si algún registro es inválido...
			continue // Lo saltamos para no romper la carga.
		}
		_ = r.mem.Agregar(*libPtr) // Agrega al repo memoria (ignora duplicados si los hay).
	}

	return nil // Carga completada.
} // Fin de cargar.

// guardar escribe el estado actual del repositorio en el archivo JSON.
func (r *RepoJSON) guardar() error { // Función para guardar a disco.
	libros := r.mem.Listar() // Obtiene todos los libros en orden.

	lista := make([]LibroDTO, 0, len(libros)) // Prepara lista DTO para JSON.
	for _, l := range libros {                // Recorre libros en memoria.
		lista = append(lista, LibroDTO{ // Convierte Libro a DTO.
			ID:     l.ID(),     // Usa getter para ID.
			Titulo: l.Titulo(), // Usa getter para Titulo.
			Autor:  l.Autor(),  // Usa getter para Autor.
			Anio:   l.Anio(),   // Usa getter para Anio.
		}) // Fin de append.
	}

	out, err := json.MarshalIndent(lista, "", "  ") // Convierte la lista a JSON bonito.
	if err != nil {                                 // Si falla el marshaling...
		return err // Devuelve error.
	}

	tmp := r.archivo + ".tmp"                            // Nombre de archivo temporal.
	if err := os.WriteFile(tmp, out, 0644); err != nil { // Escribe el JSON en tmp.
		return err // Devuelve error si no se pudo escribir.
	}

	return os.Rename(tmp, r.archivo) // Reemplaza el archivo real de forma segura.
} // Fin de guardar.

// Agregar agrega y guarda inmediatamente.
func (r *RepoJSON) Agregar(libro Libro) error { // Implementa la interfaz.
	if err := r.mem.Agregar(libro); err != nil { // Agrega en memoria.
		return err // Si falla (duplicado), devuelve error.
	}
	return r.guardar() // Guarda a disco.
} // Fin de Agregar.

// Actualizar actualiza y guarda inmediatamente.
func (r *RepoJSON) Actualizar(libro Libro) error { // Implementa la interfaz.
	if err := r.mem.Actualizar(libro); err != nil { // Actualiza en memoria.
		return err // Si falla (no existe), devuelve error.
	}
	return r.guardar() // Guarda a disco.
} // Fin de Actualizar.

// Eliminar elimina y guarda inmediatamente.
func (r *RepoJSON) Eliminar(id int) error { // Implementa la interfaz.
	if err := r.mem.Eliminar(id); err != nil { // Elimina en memoria.
		return err // Si falla (no existe), devuelve error.
	}
	return r.guardar() // Guarda a disco.
} // Fin de Eliminar.

// BuscarPorID busca en memoria (ya está cargado).
func (r *RepoJSON) BuscarPorID(id int) (Libro, error) { // Implementa la interfaz.
	return r.mem.BuscarPorID(id) // Delegación a RepoMemoria.
} // Fin de BuscarPorID.

// Listar lista desde memoria (en orden).
func (r *RepoJSON) Listar() []Libro { // Implementa la interfaz.
	return r.mem.Listar() // Delegación a RepoMemoria.
} // Fin de Listar.
